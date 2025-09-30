package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcopollivier/poc.test-pyramid/service"
)

type PriceHandler struct {
	priceService *service.PriceService
}

func NewPriceHandler(priceService *service.PriceService) *PriceHandler {
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
	var req DiscountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	finalPrice, err := ph.priceService.CalculateDiscount(req.Price, req.Discount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DiscountResponse{FinalPrice: finalPrice})
}

func (ph *PriceHandler) GetDiscount(c *gin.Context) {
	priceStr := c.Query("price")
	discountStr := c.Query("discount")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid price"})
		return
	}

	discount, err := strconv.ParseFloat(discountStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid discount"})
		return
	}

	finalPrice, err := ph.priceService.CalculateDiscount(price, discount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DiscountResponse{FinalPrice: finalPrice})
}
