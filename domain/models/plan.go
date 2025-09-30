package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Features map[string]interface{}

func (f Features) Value() (driver.Value, error) {
	return json.Marshal(f)
}

func (f *Features) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), f)
}

type SubscriptionPlan struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"not null"`
	LookupKey       string    `json:"lookup_key" gorm:"uniqueIndex"`
	Provider        string    `json:"provider" gorm:"not null"`
	ProviderPriceID string    `json:"provider_price_id" gorm:"not null"`
	Currency        string    `json:"currency" gorm:"not null"`
	Interval        string    `json:"interval" gorm:"not null"`
	Amount          int64     `json:"amount" gorm:"not null"`
	Features        Features  `json:"features" gorm:"type:jsonb"`
	Active          bool      `json:"active" gorm:"default:true"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
