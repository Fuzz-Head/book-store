package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/Fuzz-Head/domain/models"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/price"
	"github.com/stripe/stripe-go/v76/subscription"
)

type StripeService struct {
	apiKey string
}

func NewStripeService(apiKey string) *StripeService {
	stripe.Key = apiKey
	return &StripeService{apiKey: apiKey}
}

func (s *StripeService) CreateCustomer(ctx context.Context, userID uint, email, name string) (*models.Customer, error) {
	params := &stripe.CustomerParams{
		Email: stripe.String(email),
		Name:  stripe.String(name),
		Metadata: map[string]string{
			"user_id": fmt.Sprintf("%d", userID),
		},
	}

	cust, err := customer.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create Stripe customer: %w", err)
	}

	return &models.Customer{
		UserID:             userID,
		Provider:           "stripe",
		ProviderCustomerID: cust.ID,
	}, nil
}

func (s *StripeService) CreateSubscription(ctx context.Context, customerID, priceLookupKey string) (*CreateSubscriptionResponse, error) {
	prices := price.List(&stripe.PriceListParams{
		LookupKeys: stripe.StringSlice([]string{priceLookupKey}),
	})

	var targetPrice *stripe.Price
	for prices.Next() {
		targetPrice = prices.Price()
		break
	}

	if targetPrice == nil {
		return nil, errors.New("price not found for lookup key")
	}

	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customerID),
		Items: []*stripe.SubscriptionItemsParams{
			{Price: stripe.String(targetPrice.ID)},
		},
		PaymentBehavior: stripe.String("default_incomplete"),
		PaymentSettings: &stripe.SubscriptionPaymentSettingsParams{
			SaveDefaultPaymentMethod: stripe.String("on_subscription"),
		},
		Expand: stripe.StringSlice([]string{"latest_invoice.payment_intent"}),
	}

	subscription, err := subscription.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create subscription: %w", err)
	}

	return &CreateSubscriptionResponse{
		SubscriptionID: subscription.ID,
		ClientSecret:   subscription.LatestInvoice.PaymentIntent.ClientSecret,
		Status:         string(subscription.Status),
	}, nil

}

type CreateSubscriptionResponse struct {
	SubscriptionID string `json:"subscription_id"`
	ClientSecret   string `json:"client_secret"`
	Status         string `json:"status"`
}
