package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourorg/doc-agent-demo/internal/models"
)

func ListProducts(c *gin.Context) {
	products := models.GetAllProducts()
	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"count":    len(products),
	})
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")

	product, exists := models.GetProductByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// Generate simple ID
	product.ID = fmt.Sprintf("prd_%d", len(models.GetAllProducts())+1)

	models.CreateProduct(&product)

	c.JSON(http.StatusCreated, product)
}
