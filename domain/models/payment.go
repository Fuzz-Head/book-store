package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// PaymentStatus represents the status of a payment
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

// PaymentMethod represents different payment methods
type PaymentMethod string

const (
	PaymentMethodCard   PaymentMethod = "card"
	PaymentMethodUPI    PaymentMethod = "upi"
	PaymentMethodWallet PaymentMethod = "wallet"
	PaymentMethodBank   PaymentMethod = "bank_transfer"
)

// Payment represents a payment transaction
type Payment struct {
	ID                   uint                   `gorm:"primaryKey" json:"id"`
	UserID               uint                   `gorm:"not null" json:"user_id"`
	SubscriptionID       *uint                  `json:"subscription_id,omitempty"`
	Provider             string                 `gorm:"not null" json:"provider"` // "stripe" or "razorpay"
	ProviderPaymentID    string                 `gorm:"not null" json:"provider_payment_id"`
	ProviderChargeID     string                 `json:"provider_charge_id,omitempty"`
	Amount               int64                  `gorm:"not null" json:"amount"` // in smallest currency unit
	Currency             string                 `gorm:"not null" json:"currency"`
	Status               PaymentStatus          `gorm:"not null" json:"status"`
	PaymentMethod        PaymentMethod          `json:"payment_method"`
	PaymentMethodDetails PaymentMethodDetails   `gorm:"type:jsonb" json:"payment_method_details,omitempty"`
	Metadata             map[string]interface{} `gorm:"serializer:json" json:"metadata,omitempty"`
	FailureReason        string                 `json:"failure_reason,omitempty"`
	RefundedAmount       int64                  `gorm:"default:0" json:"refunded_amount"`
	ProcessedAt          *time.Time             `json:"processed_at,omitempty"`
	CreatedAt            time.Time              `json:"created_at"`
	UpdatedAt            time.Time              `json:"updated_at"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	//Invoice      *Invoice      `gorm:"foreignKey:InvoiceID" json:"invoice,omitempty"`
	Subscription *Subscription `gorm:"foreignKey:SubscriptionID" json:"subscription,omitempty"`
}

// PaymentMethodDetails stores payment method specific information
type PaymentMethodDetails struct {
	CardLast4   string `json:"card_last4,omitempty"`
	CardBrand   string `json:"card_brand,omitempty"`
	CardCountry string `json:"card_country,omitempty"`
	UPIHandle   string `json:"upi_handle,omitempty"`
	WalletType  string `json:"wallet_type,omitempty"`
	BankName    string `json:"bank_name,omitempty"`
}

// Value implements the driver.Valuer interface for PaymentMethodDetails
func (p PaymentMethodDetails) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan implements the sql.Scanner interface for PaymentMethodDetails
func (p *PaymentMethodDetails) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), p)
}

// Invoice represents billing invoices
type Invoice struct {
	ID                uint                   `gorm:"primaryKey" json:"id"`
	UserID            uint                   `gorm:"not null" json:"user_id"`
	SubscriptionID    *uint                  `json:"subscription_id,omitempty"`
	Provider          string                 `gorm:"not null" json:"provider"`
	ProviderInvoiceID string                 `gorm:"not null;uniqueIndex" json:"provider_invoice_id"`
	InvoiceNumber     string                 `json:"invoice_number,omitempty"`
	Status            InvoiceStatus          `gorm:"not null" json:"status"`
	AmountDue         int64                  `gorm:"not null" json:"amount_due"`
	AmountPaid        int64                  `gorm:"default:0" json:"amount_paid"`
	Currency          string                 `gorm:"not null" json:"currency"`
	Description       string                 `json:"description,omitempty"`
	HostedInvoiceURL  string                 `json:"hosted_invoice_url,omitempty"`
	InvoicePDF        string                 `json:"invoice_pdf,omitempty"`
	DueDate           *time.Time             `json:"due_date,omitempty"`
	PaidAt            *time.Time             `json:"paid_at,omitempty"`
	Metadata          map[string]interface{} `gorm:"serializer:json" json:"metadata,omitempty"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`

	// Relations
	User         User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Subscription *Subscription `gorm:"foreignKey:SubscriptionID" json:"subscription,omitempty"`
	//Payments     []Payment     `gorm:"foreignKey:PaymentID" json:"payments,omitempty"`
}

// InvoiceStatus represents the status of an invoice
type InvoiceStatus string

const (
	InvoiceStatusDraft         InvoiceStatus = "draft"
	InvoiceStatusOpen          InvoiceStatus = "open"
	InvoiceStatusPaid          InvoiceStatus = "paid"
	InvoiceStatusVoid          InvoiceStatus = "void"
	InvoiceStatusUncollectible InvoiceStatus = "uncollectible"
)

// PaymentIntent represents payment intents for subscription setup
type PaymentIntent struct {
	ID                 uint                   `gorm:"primaryKey" json:"id"`
	UserID             uint                   `gorm:"not null" json:"user_id"`
	Provider           string                 `gorm:"not null" json:"provider"`
	ProviderIntentID   string                 `gorm:"not null;uniqueIndex" json:"provider_intent_id"`
	Amount             int64                  `gorm:"not null" json:"amount"`
	Currency           string                 `gorm:"not null" json:"currency"`
	Status             PaymentIntentStatus    `gorm:"not null" json:"status"`
	ClientSecret       string                 `json:"client_secret,omitempty"`
	ConfirmationMethod string                 `json:"confirmation_method,omitempty"`
	PaymentMethodTypes []string               `gorm:"serializer:json" json:"payment_method_types,omitempty"`
	Metadata           map[string]interface{} `gorm:"serializer:json" json:"metadata,omitempty"`
	LastPaymentError   string                 `json:"last_payment_error,omitempty"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// PaymentIntentStatus represents the status of a payment intent
type PaymentIntentStatus string

const (
	PaymentIntentStatusRequiresPaymentMethod PaymentIntentStatus = "requires_payment_method"
	PaymentIntentStatusRequiresConfirmation  PaymentIntentStatus = "requires_confirmation"
	PaymentIntentStatusRequiresAction        PaymentIntentStatus = "requires_action"
	PaymentIntentStatusProcessing            PaymentIntentStatus = "processing"
	PaymentIntentStatusSucceeded             PaymentIntentStatus = "succeeded"
	PaymentIntentStatusCanceled              PaymentIntentStatus = "canceled"
)

// Refund represents payment refunds
type Refund struct {
	ID               uint                   `gorm:"primaryKey" json:"id"`
	PaymentID        uint                   `gorm:"not null" json:"payment_id"`
	Provider         string                 `gorm:"not null" json:"provider"`
	ProviderRefundID string                 `gorm:"not null;uniqueIndex" json:"provider_refund_id"`
	Amount           int64                  `gorm:"not null" json:"amount"`
	Currency         string                 `gorm:"not null" json:"currency"`
	Status           RefundStatus           `gorm:"not null" json:"status"`
	Reason           string                 `json:"reason,omitempty"`
	Metadata         map[string]interface{} `gorm:"serializer:json" json:"metadata,omitempty"`
	ProcessedAt      *time.Time             `json:"processed_at,omitempty"`
	CreatedAt        time.Time              `json:"created_at"`
	UpdatedAt        time.Time              `json:"updated_at"`

	// Relations
	Payment Payment `gorm:"foreignKey:PaymentID" json:"payment"`
}

// RefundStatus represents the status of a refund
type RefundStatus string

const (
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusSucceeded RefundStatus = "succeeded"
	RefundStatusFailed    RefundStatus = "failed"
	RefundStatusCanceled  RefundStatus = "canceled"
)

// PaymentAnalytics represents payment analytics data
type PaymentAnalytics struct {
	TotalRevenue       int64               `json:"total_revenue"`
	MonthlyRevenue     int64               `json:"monthly_revenue"`
	TotalTransactions  int64               `json:"total_transactions"`
	SuccessfulPayments int64               `json:"successful_payments"`
	FailedPayments     int64               `json:"failed_payments"`
	RefundedAmount     int64               `json:"refunded_amount"`
	AverageOrderValue  float64             `json:"average_order_value"`
	PaymentMethodStats map[string]int64    `json:"payment_method_stats"`
	RevenueByMonth     map[string]int64    `json:"revenue_by_month"`
	TopPaymentMethods  []PaymentMethodStat `json:"top_payment_methods"`
}

// PaymentMethodStat represents statistics for a payment method
type PaymentMethodStat struct {
	Method      PaymentMethod `json:"method"`
	Count       int64         `json:"count"`
	TotalAmount int64         `json:"total_amount"`
	Percentage  float64       `json:"percentage"`
}

// WebhookEvent represents incoming webhook events from payment providers
type WebhookEvent struct {
	ID              uint                   `gorm:"primaryKey" json:"id"`
	Provider        string                 `gorm:"not null" json:"provider"`
	EventType       string                 `gorm:"not null" json:"event_type"`
	ProviderEventID string                 `gorm:"uniqueIndex" json:"provider_event_id"`
	Data            map[string]interface{} `gorm:"serializer:json" json:"data"`
	Processed       bool                   `gorm:"default:false" json:"processed"`
	ProcessedAt     *time.Time             `json:"processed_at,omitempty"`
	ErrorMessage    string                 `json:"error_message,omitempty"`
	RetryCount      int                    `gorm:"default:0" json:"retry_count"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

// IsSuccessful checks if payment status indicates success
func (p *Payment) IsSuccessful() bool {
	return p.Status == PaymentStatusCompleted
}

// IsFailed checks if payment status indicates failure
func (p *Payment) IsFailed() bool {
	return p.Status == PaymentStatusFailed || p.Status == PaymentStatusCancelled
}

// IsRefundable checks if payment can be refunded
func (p *Payment) IsRefundable() bool {
	return p.Status == PaymentStatusCompleted && p.RefundedAmount < p.Amount
}

// GetRefundableAmount returns the amount that can still be refunded
func (p *Payment) GetRefundableAmount() int64 {
	return p.Amount - p.RefundedAmount
}

// FormatAmount formats the amount in a readable format (e.g., $10.99)
func (p *Payment) FormatAmount() string {
	switch p.Currency {
	case "USD":
		return fmt.Sprintf("$%.2f", float64(p.Amount)/100)
	case "INR":
		return fmt.Sprintf("₹%.2f", float64(p.Amount)/100)
	case "EUR":
		return fmt.Sprintf("€%.2f", float64(p.Amount)/100)
	default:
		return fmt.Sprintf("%.2f %s", float64(p.Amount)/100, p.Currency)
	}
}

// IsOverdue checks if invoice is overdue
func (i *Invoice) IsOverdue() bool {
	if i.DueDate == nil || i.Status == InvoiceStatusPaid {
		return false
	}
	return time.Now().After(*i.DueDate)
}

// GetOutstandingAmount returns the amount still owed on the invoice
func (i *Invoice) GetOutstandingAmount() int64 {
	return i.AmountDue - i.AmountPaid
}
