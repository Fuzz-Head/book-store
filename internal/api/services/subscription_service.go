package services

import (
	"context"
	"errors"

	"github.com/Fuzz-Head/domain/models"
	"gorm.io/gorm"
)

type SubscriptionService struct {
	db              *gorm.DB
	stripeService   *StripeService
	razorpayService *RazorpayService
}

func NewSubscriptionService(db *gorm.DB, stripeService *StripeService, razorpayService *RazorpayService) *SubscriptionService {
	return &SubscriptionService{
		db:              db,
		stripeService:   stripeService,
		razorpayService: razorpayService,
	}
}

func (s *SubscriptionService) GetPlans(ctx context.Context) ([]models.SubscriptionPlan, error) {
	var plans []models.SubscriptionPlan
	err := s.db.Where("active = ?", true).Find(&plans).Error
	return plans, err
}

func (s *SubscriptionService) CreateCheckoutSession(ctx context.Context, userID uint, planLookupKey string) (*CheckoutResponse, error) {
	var plan models.SubscriptionPlan
	err := s.db.Where("lookup_key = ? AND active = ?", planLookupKey, true).First(&plan).Error
	if err != nil {
		return nil, errors.New("plan not found")
	}

	customer, err := s.GetOrCreateCustomer(ctx, userID, plan.Provider)
	if err != nil {
		return nil, err
	}

	var response *CreateSubscriptionResponse
	switch plan.Provider {
	case "stripe":
		response, err = s.stripeService.CreateSubscription(ctx, customer.ProviderCustomerID, planLookupKey)
	case "razorpay":
		response, err = s.razorpayService.CreateSubscription(ctx, plan.ProviderPriceID)
	default:
		return nil, errors.New("unsupported payment provider")
	}

	if err != nil {
		return nil, err
	}

	subscription := models.Subscription{
		UserID:                 userID,
		PlanID:                 plan.ID,
		Provider:               plan.Provider,
		ProviderSubscriptionID: response.SubscriptionID,
		Status:                 models.SubscriptionStatus(response.Status),
	}

	err = s.db.Create(&subscription).Error
	if err != nil {
		return nil, err
	}

	return &CheckoutResponse{
		Provider:       plan.Provider,
		SubscriptionID: response.SubscriptionID,
		ClientSecret:   response.ClientSecret,
		Status:         response.Status,
	}, nil
}

func (s *SubscriptionService) GetOrCreateCustomer(ctx context.Context, userID uint, provider string) (*models.Customer, error) {
	var customer models.Customer
	err := s.db.Where("user_id = ? AND provider = ?", userID, provider).First(&customer).Error
	if err == nil {
		return &customer, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// get values from user service
	email := "user@example.com"
	name := "User Name"

	switch provider {
	case "stripe":
		return s.stripeService.CreateCustomer(ctx, userID, email, name)
	case "razorpay":
		return s.razorpayService.CreateCustomer(ctx, userID, email, name)
	default:
		return nil, errors.New("unsupported provider")
	}
}

func (s *SubscriptionService) GetUserSubscriptions(ctx context.Context, userID uint) ([]models.Subscription, error) {
	var subscription []models.Subscription
	err := s.db.Where("user_id = ?", userID).Find(&subscription).Error
	return subscription, err
}

func (s *SubscriptionService) GetSubscriptions(ctx context.Context, subscription_ID string, userID uint) (*[]models.Subscription, error) {
	var subscription []models.Subscription
	err := s.db.Preload("Plan").Where("provider_subscription_id = ? AND user_id = ?", userID, subscription_ID).Find(&subscription).Error
	return &subscription, err
}

func (s *SubscriptionService) CancelSubscription(ctx context.Context, subscriptionID string, userID uint) error {
	var subscription models.Subscription
	err := s.db.Where("provider_subscription_id = ? AND user_id = ?", subscriptionID, userID).First(&subscription).Error
	if err != nil {
		return err
	}

	subscription.Status = models.StatusCancelled
	return s.db.Save(&subscription).Error
}

func (s *SubscriptionService) UpdateSubscription(ctx context.Context, subscriptionID string, userID uint, newPlanLookupKey string, cancelAtPeriodEnd *bool) (*models.Subscription, error) {
	var subscription models.Subscription
	err := s.db.Where("provider_subscription_id = ? AND user_id = ?", subscriptionID, userID).First(&subscription).Error
	if err != nil {
		return nil, err
	}

	// Update cancel_at_period_end if provided
	if cancelAtPeriodEnd != nil {
		subscription.CancelAtPeriodEnd = *cancelAtPeriodEnd
		s.db.Save(&subscription)
	}

	// TODO: Implement plan changes with provider APIs
	return &subscription, nil
}

type CheckoutResponse struct {
	Provider       string `json:"provider"`
	SubscriptionID string `json:"subscription_id"`
	ClientSecret   string `json:"client_secret,omitempty"`
	Status         string `json:"status"`
}
