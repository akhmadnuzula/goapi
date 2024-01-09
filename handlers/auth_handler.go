package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goapi/models"
	"goapi/database"

	// "github.com/swaggo/swag/example/celler/httputil"
)

// get login by id
// @Summary Get a login by ID
// @Description Get a login by ID
// @Tags login
// @Accept  json
// @Produce  json
// @Param id path int true "Login ID"
// @Success 200 {object} gin.H
// @Router /login/{id} [get]
func Login(c *gin.Context) {
	var loginInfo struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data login tidak valid" . loginInfo.Username})
		return
	}

	var user models.User
	// cari user
	if err:= database.DB.Where("username = ?", loginInfo.Username).First(&user); err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "akun tidak ditemukan"})
    return
	}
	
	// cek password
	if user.Password != loginInfo.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password anda salah"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}