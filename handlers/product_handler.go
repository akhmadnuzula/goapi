package handlers

import (
	"net/http"
	"strconv"

	"goapi/database"
	"goapi/models"

	"github.com/gin-gonic/gin"
	// "github.com/swaggo/swag/example/celler/httputil"
)

// get all products
// @Summary Get list of product
// @Description Get list of product
// @Tags products
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// get products by id
// @Summary Get a product by ID
// @Description Get a product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /products/{id} [get]
func GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	database.DB.First(&product, id)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// create products
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept  json
// @Produce  json
// @Param input body models.Product true "Product input"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var productInput models.Product
	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	database.DB.Create(&productInput)
	c.JSON(http.StatusOK, gin.H{"data": productInput})
}

// update product by id
// @Summary Update a product by ID
// @Description Update a product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param input body models.Product true "Product input"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var productInput models.Product
	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	database.DB.First(&product, id)
	product.Name = productInput.Name
	product.Price = productInput.Price
	product.Stock = productInput.Stock
	database.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// delete product
// @Summary Delete a product by ID
// @Description Delete a product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Delete(&models.Product{}, id)

	c.JSON(http.StatusOK, gin.H{"data": "Product Deleted"})
}