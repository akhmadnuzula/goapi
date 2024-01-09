package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"goapi/models"
	"goapi/database"

	// "github.com/swaggo/swag/example/celler/httputil"
)

// get all product
// @Summary Get list of product
// @Description Get list of product
// @Tags product
// @Accept  json
// @Produce  json
// @Success 200 {object} gin.H
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
// @Success 200 {object} gin.H
// @Router /products/{id} [get]
func GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	database.DB.First(&product, id)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// create product
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept  json
// @Produce  json
// @Param input body models.Product true "Product input"
// @Success 200 {object} gin.H
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
// @Tags product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param input body models.Product true "Product input"
// @Success 200 {object} gin.H
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
// @Success 200 {object} gin.H
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Delete(&models.Product{}, id)

	c.JSON(http.StatusOK, gin.H{"data": "Product Deleted"})
}