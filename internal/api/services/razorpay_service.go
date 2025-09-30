package services

import (
	"context"
	"fmt"

	"github.com/Fuzz-Head/domain/models"
	razorpay "github.com/razorpay/razorpay-go"
)

type RazorpayService struct {
	client *razorpay.Client
}

func NewRazorpayService(keyID, keySecret string) *RazorpayService {
	client := razorpay.NewClient(keyID, keySecret)
	return &RazorpayService{client: client}
}

func (r *RazorpayService) CreateCustomer(ctx context.Context, userID uint, email, name string) (*models.Customer, error) {
	data := map[string]interface{}{
		"name":  name,
		"email": email,
		"notes": map[string]interface{}{
			"user_id": fmt.Sprintf("%d", userID),
		},
	}

	cust, err := r.client.Customer.Create(data, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Razorpay customer: %w", err)
	}

	return &models.Customer{
		UserID:             userID,
		Provider:           "razorpay",
		ProviderCustomerID: cust["id"].(string),
	}, nil
}

func (r *RazorpayService) CreateSubscription(ctx context.Context, planID string) (*CreateSubscriptionResponse, error) {
	data := map[string]interface{}{
		"plan_id": planID,
		"notes": map[string]interface{}{
			"source": "bookstore_api",
		},
	}

	sub, err := r.client.Subscription.Create(data, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create a Subscription: %w", err)
	}

	return &CreateSubscriptionResponse{
		SubscriptionID: sub["id"].(string),
		Status:         sub["status"].(string),
	}, nil
}
