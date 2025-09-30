package models

import "time"

type SubscriptionStatus string

const (
	StatusActive     SubscriptionStatus = "active"
	StatusPaused     SubscriptionStatus = "paused"
	StatusCancelled  SubscriptionStatus = "cancelled"
	StatusIncomplete SubscriptionStatus = "incomplete"
	StatusPastDue    SubscriptionStatus = "past_due"
	StatusTrialing   SubscriptionStatus = "trialing"
)

type Customer struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	UserID               uint      `gorm:"not null;uniqueIndex:idx_user_provider" json:"user_id"`
	Provider             string    `gorm:"not null;uniqueIndex:idx_user_provider" json:"provider"`
	ProviderCustomerID   string    `gorm:"not null" json:"provider_customer_id"`
	DefaultPaymentMethod string    `json:"default_payment_method,omitempty"`
	CreatedAt            time.Time `json:"created_at"`
}

type Subscription struct {
	ID                     uint               `gorm:"primaryKey" json:"id"`
	UserID                 uint               `gorm:"not null" json:"user_id"`
	PlanID                 uint               `gorm:"not null" json:"plan_id"`
	Provider               string             `gorm:"not null" json:"provider"`
	ProviderSubscriptionID string             `gorm:"not null;uniqueIndex" json:"provider_subscription_id"`
	Status                 SubscriptionStatus `gorm:"not null" json:"status"`
	CurrentPeriodStart     *time.Time         `json:"current_period_start"`
	CurrentPeriodEnd       *time.Time         `json:"current_period_end"`
	CancelAtPeriodEnd      bool               `gorm:"default:false" json:"cancel_at_period_end"`
	TrialEnd               *time.Time         `json:"trial_end,omitempty"`
	CreatedAt              time.Time          `json:"created_at"`
	UpdatedAt              time.Time          `json:"updated_at"`

	// Relations
	Plan     SubscriptionPlan `gorm:"foreignKey:PlanID" json:"plan"`
	Customer Customer         `gorm:"foreignKey:UserID;references:UserID" json:"customer"`
}
