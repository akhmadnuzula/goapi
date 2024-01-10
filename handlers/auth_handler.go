package handlers

import (
	"net/http"
	"os"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data login tidak valid"})
		return
	}

	var users []models.User
	// cari user
	database.DB.Where("username = ?", loginInfo.Username).Find(&users)
	
	if len(users) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Akun tidak ditemukan"})
		return
	}

	// cek password
	for _, user := range users {
		if user.Password == loginInfo.Password {
			token := os.Getenv("TOKEN")
			c.JSON(http.StatusOK, gin.H{"token": token})
			return
		}
	}

	// password tidak di temukan
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Password Anda salah"})
}
