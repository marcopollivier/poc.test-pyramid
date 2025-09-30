package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcopollivier/poc.test-pyramid/service"
)

type PriceHandler struct {
	priceService *service.PriceService
}

func NewPriceHandler(priceService *service.PriceService) *PriceHandler {
	log.Println("PriceHandler: Initializing handler")
	return &PriceHandler{priceService: priceService}
}

type DiscountRequest struct {
	Price    float64 `json:"price" binding:"required"`
	Discount float64 `json:"discount" binding:"required"`
}

type DiscountResponse struct {
	FinalPrice float64 `json:"final_price"`
}

func (ph *PriceHandler) CalculateDiscount(c *gin.Context) {
	log.Println("PriceHandler: POST /discount - Request received")
	
	var req DiscountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("PriceHandler: JSON binding error - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("PriceHandler: Request data - Price: %.2f, Discount: %.2f%%", req.Price, req.Discount)

	finalPrice, err := ph.priceService.CalculateDiscount(req.Price, req.Discount)
	if err != nil {
		log.Printf("PriceHandler: Service error - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := DiscountResponse{FinalPrice: finalPrice}
	log.Printf("PriceHandler: Sending response - Final price: %.2f", finalPrice)
	c.JSON(http.StatusOK, response)
}

func (ph *PriceHandler) GetDiscount(c *gin.Context) {
	log.Println("PriceHandler: GET /discount - Request received")
	
	priceStr := c.Query("price")
	discountStr := c.Query("discount")
	
	log.Printf("PriceHandler: Query params - price: %s, discount: %s", priceStr, discountStr)

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		log.Printf("PriceHandler: Price parsing error - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid price"})
		return
	}

	discount, err := strconv.ParseFloat(discountStr, 64)
	if err != nil {
		log.Printf("PriceHandler: Discount parsing error - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid discount"})
		return
	}

	log.Printf("PriceHandler: Parsed data - Price: %.2f, Discount: %.2f%%", price, discount)

	finalPrice, err := ph.priceService.CalculateDiscount(price, discount)
	if err != nil {
		log.Printf("PriceHandler: Service error - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := DiscountResponse{FinalPrice: finalPrice}
	log.Printf("PriceHandler: Sending response - Final price: %.2f", finalPrice)
	c.JSON(http.StatusOK, response)
}
