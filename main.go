package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcopollivier/poc.test-pyramid/handler"
	"github.com/marcopollivier/poc.test-pyramid/service"
)

func main() {
	r := gin.Default()

	priceService := service.NewPriceService()
	priceHandler := handler.NewPriceHandler(priceService)

	r.POST("/discount", priceHandler.CalculateDiscount)
	r.GET("/discount", priceHandler.GetDiscount)

	r.Run(":8080")
}
