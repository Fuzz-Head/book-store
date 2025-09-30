package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Fuzz-Head/database"
	"github.com/Fuzz-Head/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/webhook"
)

func StripeWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error reading Stripe webhook body: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "error reading request"})
		return
	}

	// verify webhook signature
	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	if endpointSecret == "" {
		log.Printf("STRIPE_WEBHOOK_SECRET not set")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Webhook secret key not configured"})
		return
	}

	event, err := webhook.ConstructEvent(payload, c.GetHeader("Stripe-Signature"), endpointSecret)
	if err != nil {
		log.Printf("Error verifying Stripe webhook signature: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Signature"})
		return
	}

	// Log webhook event for debugging
	log.Printf("Received Stripe webhook event: %s (ID: %s)", event.Type, event.ID)

	// Store webhook event in database for audit trial
	webhookEvent := models.WebhookEvent{
		Provider:        "stripe",
		EventType:       string(event.Type),
		ProviderEventID: event.ID,
		Data:            make(map[string]interface{}),
		Processed:       false,
	}

	// Convert event data to map for Store
	eventDataBytes, _ := json.Marshal(event.Data.Object)
	json.Unmarshal(eventDataBytes, &webhookEvent.Data)

	if err := database.DB.Create(&webhookEvent).Error; err != nil {
		log.Printf("Failed to store webhook event: %v", err)
	}

	// Process the event
	switch event.Type {
	case "chcekout.session.completed":
		err = handleStripeCheckoutCompleted(event)
	case "invoice.paid":
		err = handleStripeInvoicePaid(event)
	case "invoice.payment_failed":
		err = handleStripeInvoicePaymentFailed(event)
	case "customer.subscription.created":
		err = handleStripeSubscriptionCreated(event)
	case "customer.subscription.updated":
		err = handleStripeSubscriptionUpdated(event)
	case "customer.subscription.deleted":
		err = handleStripeSubscriptionDeleted(event)
	case "payment_intent.succeeded":
		err = handleStripePaymentSucceeded(event)
	case "payment_intent.payment_failed":
		err = handleStripePaymentFailed(event)
	default:
		log.Printf("Unhandled Stripe event type: %s", event.Type)
		err = nil
	}

	// Update webhook processing Status
	if err != nil {
		log.Printf("Error processing Stripe webhook %s: %v", event.Type, err)
		webhookEvent.ErrorMessage = err.Error()
		webhookEvent.RetryCount++
	} else {
		webhookEvent.Processed = true
		now := time.Now()
		webhookEvent.ProcessedAt = &now
	}

	database.DB.Save(&webhookEvent)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process webhook"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

// Stripe event handlers
func handleStripeCheckoutCompleted(event stripe.Event) error {
	var session stripe.CheckoutSession
	if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
		return fmt.Errorf("error parsing checkout session: %v", err)
	}

	log.Printf("Checout session completed: %s", session.ID)

	// update subscription if subscription checkout
	if session.Mode == "subscription" && session.Subscription != nil {
		return UpdateSubscriptionFromStripe(session.Subscription.ID, models.StatusActive)
	}

	return nil
}

func handleStripeInvoicePaid(event stripe.Event) error {
	var invoice stripe.Invoice
	if err := json.Unmarshal(event.Data.Raw, &invoice); err != nil {
		return fmt.Errorf("error parsing invoice: %v", err)
	}

	log.Printf("invoice paid: %s", invoice.ID)

	// Update subscription status
	if invoice.Subscription != nil {
		return UpdateSubscriptionFromStripe(invoice.Subscription.ID, models.StatusActive)
	}
	return nil
}

func handleStripeInvoicePaymentFailed(event stripe.Event) error {
	var invoice stripe.Invoice
	if err := json.Unmarshal(event.Data.Raw, &invoice); err != nil {
		return fmt.Errorf("error parsing invoice: %v", err)
	}

	log.Printf("Invoice payment failed: %s", invoice.ID)

	if invoice.Subscription != nil {
		return UpdateSubscriptionFromStripe(invoice.Subscription.ID, models.StatusPastDue)
	}

	return nil
}

func handleStripeSubscriptionCreated(event stripe.Event) error {
	var subscription stripe.Subscription
	if err := json.Unmarshal(event.Data.Raw, &subscription); err != nil {
		return fmt.Errorf("error parsing subscription: %v", err)
	}

	log.Printf("Subscription created: %s", subscription.ID)
	return UpdateSubscriptionFromStripe(subscription.ID, models.SubscriptionStatus(subscription.Status))
}

func handleStripeSubscriptionUpdated(event stripe.Event) error {
	var subscription stripe.Subscription
	if err := json.Unmarshal(event.Data.Raw, &subscription); err != nil {
		return fmt.Errorf("error parsing subscription: %v", err)
	}

	log.Printf("Subscription updated: %s", subscription.ID)
	return UpdateSubscriptionFromStripe(subscription.ID, models.SubscriptionStatus(subscription.Status))
}

func handleStripeSubscriptionDeleted(event stripe.Event) error {
	var subscription stripe.Subscription
	if err := json.Unmarshal(event.Data.Raw, &subscription); err != nil {
		return fmt.Errorf("error parsing subscription: %v", err)
	}

	log.Printf("Subscription deleted: %s", subscription.ID)
	return UpdateSubscriptionFromStripe(subscription.ID, models.StatusCancelled)
}

func handleStripePaymentSucceeded(event stripe.Event) error {
	var paymentIntent stripe.PaymentIntent
	if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
		return fmt.Errorf("error parsing payment intent: %v", err)
	}

	log.Printf("Payment failed: %s", paymentIntent.ID)
	// TODO: Create payment record
	return nil
}

func handleStripePaymentFailed(event stripe.Event) error {
	var paymentIntent stripe.PaymentIntent
	if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
		return fmt.Errorf("error parsing payment intent: %v", err)
	}

	log.Printf("Payment failed: %s", paymentIntent.ID)
	// TODO: Handle Failed payments
	return nil
}

// TODO: Handle all of Razorpay

// Helper function

func UpdateSubscriptionFromStripe(stripeSubscriptionID string, status models.SubscriptionStatus) error {
	var subscription models.Subscription
	err := database.DB.Where("provide_subscription_id = ? AND provider =?", stripeSubscriptionID, "stripe").First(&subscription).Error
	if err != nil {
		return fmt.Errorf("subscription not found: %v", err)
	}

	subscription.Status = status
	return database.DB.Save(&subscription).Error
}
