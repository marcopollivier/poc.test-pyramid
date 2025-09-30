package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/marcopollivier/poc.test-pyramid/config"
	"github.com/marcopollivier/poc.test-pyramid/handler"
	"github.com/marcopollivier/poc.test-pyramid/messaging"
	"github.com/marcopollivier/poc.test-pyramid/repository"
	"github.com/marcopollivier/poc.test-pyramid/service"
)

func main() {
	log.Println("=== POC Test Pyramid Starting ===")
	
	// Setup database
	log.Println("Main: Setting up database connection...")
	db := config.SetupDatabase()

	// Setup Kafka
	log.Println("Main: Setting up Kafka publisher...")
	kafkaPublisher, err := messaging.NewKafkaPublisher([]string{"localhost:9092"})
	if err != nil {
		log.Fatal("Main: Failed to create Kafka publisher:", err)
	}
	defer func() {
		log.Println("Main: Closing Kafka publisher...")
		kafkaPublisher.Close()
	}()

	// Setup dependencies
	log.Println("Main: Initializing application components...")
	discountRepo := repository.NewDiscountRepository(db)
	priceService := service.NewPriceService(discountRepo, kafkaPublisher)
	priceHandler := handler.NewPriceHandler(priceService)

	// Setup routes
	log.Println("Main: Setting up HTTP routes...")
	r := gin.Default()
	r.POST("/discount", priceHandler.CalculateDiscount)
	r.GET("/discount", priceHandler.GetDiscount)

	log.Println("=== Server starting on :8080 ===")
	log.Println("API Endpoints:")
	log.Println("  POST http://localhost:8080/discount")
	log.Println("  GET  http://localhost:8080/discount")
	log.Println("Kafka UI: http://localhost:8090")
	log.Println("=====================================")
	
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Main: Failed to start server:", err)
	}
}
