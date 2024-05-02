package main

import (
	"github.com/gin-gonic/gin"

	"Delivery_tracking_system/handlers"
)

func main() {
	r := gin.Default()

	//routes for user
	r.POST("/api/register", handlers.Register)
	r.POST("/api/login", handlers.Login)
	r.GET("/api/profile", handlers.Profile)

	//routes for package managements
	r.POST("/api/packages", handlers.CreatePackage)
	r.GET("/api/packages/:id", handlers.GetPackage)
	r.PUT("/api/packages/:id", handlers.UpdatePackage)
	r.DELETE("/api/packages/:id", handlers.DeletePackage)

	//routes for location tracking
	r.POST("api/locations/update", handlers.UpdateLocation)
	r.GET("api/locations/:id", handlers.GetLocation)
	r.GET("api/locations/history/:id", handlers.GetLocationHistory)
	r.GET("api/locations/search", handlers.SearchLocation)
	r.POST("api/geo-fencing/check", handlers.CheckGeoFencing)

	//routes for notification and alerts
	r.POST("api/notifications", handlers.SendNotification)
	r.POST("api/alerts", handlers.SetAlerts)

	//routes for order tracking
	r.GET("/api/orders/:trackingNumber", handlers.TrackingOrder)
	r.GET("/api/orders/:trackingNumber/history", handlers.GetOrderHistory)

	//routes for intergration with other system
	r.POST("/api/integrations/orders", handlers.CreateOrderIntegration)
	r.POST("/api/webhooks/:event", handlers.WebHookHandler)

	//routes for reporting and analytics
	r.GET("/api/analytics", handlers.GetAnalytics)
}
