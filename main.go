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
	// Setup database
	db := config.SetupDatabase()

	// Setup Kafka
	kafkaPublisher, err := messaging.NewKafkaPublisher([]string{"localhost:9092"})
	if err != nil {
		log.Fatal("Failed to create Kafka publisher:", err)
	}
	defer kafkaPublisher.Close()

	// Setup dependencies
	discountRepo := repository.NewDiscountRepository(db)
	priceService := service.NewPriceService(discountRepo, kafkaPublisher)
	priceHandler := handler.NewPriceHandler(priceService)

	// Setup routes
	r := gin.Default()
	r.POST("/discount", priceHandler.CalculateDiscount)
	r.GET("/discount", priceHandler.GetDiscount)

	log.Println("Server starting on :8080")
	r.Run(":8080")
}
