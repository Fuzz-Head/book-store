package handlers

import (
	"net/http"
	"strconv"

	"github.com/Fuzz-Head/database"
	"github.com/Fuzz-Head/domain/models"
	"github.com/Fuzz-Head/internal/api/middleware"
	"github.com/Fuzz-Head/internal/api/services"
	"github.com/gin-gonic/gin"
)

var subscriptionService *services.SubscriptionService

// SetSubscriptionService allows setting the serivce dependency
func SetSubscriptionService(service *services.SubscriptionService) {
	subscriptionService = service
}

func GetSubscriptionPlans(c *gin.Context) {
	var plans []models.SubscriptionPlan

	if err := database.DB.Where("active = ?", true).Find(&plans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subscription plans"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": plans})
}

func CreateOrGetCustomer(c *gin.Context) {
	var req struct {
		Provider string `json:"provider" binding:"required,oneof=stripe razorpay"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, ok := middleware.GetClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, err := strconv.ParseUint(claims.UserID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var customer models.Customer
	err = database.DB.Where("user_id = ? AND provider = ?", userID, req.Provider).Error

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"data": customer})
		return
	}

	customer = models.Customer{
		UserID:             uint(userID),
		Provider:           req.Provider,
		ProviderCustomerID: "temp_customer_" + strconv.Itoa(int(userID)),
	}

	if err := database.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func CreateCheckoutSession(c *gin.Context) {
	if subscriptionService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Subscription services are not intiated yet"})
		return
	}

	var req struct {
		PlanLookupKey string `json:"plan_lookup_key" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, ok := middleware.GetClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, err := strconv.ParseUint(claims.UserID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	response, err := subscriptionService.CreateCheckoutSession(c.Request.Context(), uint(userID), req.PlanLookupKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func GetUserSubscriptions(c *gin.Context) {
	if subscriptionService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Subscription service not initialised"})
		return
	}

	claims, ok := middleware.GetClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authentication"})
		return
	}

	userID, err := strconv.ParseUint(claims.UserID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	subscriptions, err := subscriptionService.GetUserSubscriptions(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subscriptions})
}

func GetSubscription(c *gin.Context) {
	if subscriptionService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Subscription serivce not initialised"})
		return
	}

	subscriptionID := c.Param("id")

	claims, ok := middleware.GetClaims(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not authenticated"})
		return
	}

	userID, err := strconv.ParseUint(claims.UserID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	subscription, err := subscriptionService.GetSubscriptions(c.Request.Context(), subscriptionID, uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subscription})
}

func UpdateSubscription(c *gin.Context) {
	if subscriptionService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Subscription service not initialised"})
		return
	}

	subscriptionID := c.Param("id")

	claims, ok := middleware.GetClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, err := strconv.ParseUint(claims.UserID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var req struct {
		PlanLookupKey     string `json:"plan_lookup_key,omitempty"`
		CancelAtPeriodEnd *bool  `json:"cancel_at_period_end,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscription, err := subscriptionService.UpdateSubscription(c.Request.Context(), subscriptionID, uint(userID), req.PlanLookupKey, req.CancelAtPeriodEnd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subscription})
}

func CancelSubscription(c *gin.Context) {
	if subscriptionService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Subscription service not initialised"})
		return
	}

	subscriptionID := c.Param("id")

	claims, ok := middleware.GetClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
	}

	userID, err := strconv.ParseUint(claims.UserID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
	}

	err = subscriptionService.CancelSubscription(c.Request.Context(), subscriptionID, uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subscription cnacelled successfully"})
}

func PauseSubscription(c *gin.Context) {
	subscriptionID := c.Param("id")

	claims, ok := middleware.GetClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, err := strconv.ParseUint(claims.UserID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var subscription models.Subscription
	err = database.DB.Where("provider_subscription_id = ? AND user_id = ?", subscriptionID, userID).First(&subscription).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}

	subscription.Status = models.StatusPaused

	database.DB.Save(&subscription)

	c.JSON(http.StatusOK, gin.H{"message": "subscription pasued successfully"})

}

func ResumeSubscription(c *gin.Context) {
	subscriptionID := c.Param("id")

	claims, ok := middleware.GetClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, err := strconv.ParseUint(claims.UserID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	// Find and update subscription status
	var subscription models.Subscription
	err = database.DB.Where("provider_subscription_id = ? AND user_id = ?", subscriptionID, userID).First(&subscription).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}

	subscription.Status = models.StatusPaused
	database.DB.Save(&subscription)

	c.JSON(http.StatusNotFound, gin.H{"error": "Subscription paused successfully"})
}

func GetUserInvoices(c *gin.Context) {
	claims, ok := middleware.GetClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// TODO: Implement invoices fetching from payment providers
	c.JSON(http.StatusOK, gin.H{
		"data":    []interface{}{},
		"message": "Invoices endpoint - to be implemented",
		"user_id": claims.UserID,
	})
}

func GetUsageStats(c *gin.Context) {
	claims, ok := middleware.GetClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// TODO: Implement usage statistics calculation
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"books_read":                  0,
			"reading_time":                0,
			"subscription_days_remaining": 0,
		},
		"message": "Usage stats endpoint - to be implemented",
		"user_id": claims.UserID,
	})
}

func GetAllSubscriptions(c *gin.Context) {
	var subscriptions []models.Subscription

	err := database.DB.Preload("Plan").Find(&subscriptions).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subscription information"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subscriptions})
}

func GetBillingAnalytics(c *gin.Context) {
	//TODO: implement actual analytics calculations
	analytics := map[string]interface{}{
		"total_subscriptions":  0,
		"active_subscriptions": 0,
		"monthly_revenue":      0,
		"churn_rate":           0,
		"most_popular_plan":    "basic-monthly",
	}

	var totalCount, activeCount int64
	database.DB.Model(&models.Subscription{}).Count(&totalCount)
	database.DB.Model(&models.Subscription{}).Where("status = ?", models.StatusActive).Count(&activeCount)

	analytics["total_subscriptions"] = totalCount
	analytics["active_subscriptions"] = activeCount

	c.JSON(http.StatusOK, gin.H{"data": analytics})
}

func CreateSubscriptionPlan(c *gin.Context) {
	var plan models.SubscriptionPlan

	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&plan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a subscription plan"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": plan})
}

func UpdateSubscriptionPlan(c *gin.Context) {
	planID := c.Param("id")

	var plan models.SubscriptionPlan
	if err := database.DB.First(&plan, planID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		return
	}

	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&plan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update plan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": plan})
}

func DeactivateScbscriptionPlan(c *gin.Context) {
	planID := c.Param("id")

	var plan models.SubscriptionPlan
	if err := database.DB.First(&plan, planID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		return
	}

	plan.Active = false
	if err := database.DB.Save(&plan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deactivate plan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plan deactivated successfully"})
}

func ProcessRefund(c *gin.Context) {
	invoiceID := c.Param("invoice_id")

	// TODO: implement actual refund processing
	c.JSON(http.StatusOK, gin.H{
		"message":    "Refund processed successfully",
		"invoice_id": invoiceID,
		"status":     "refunded",
	})
}
