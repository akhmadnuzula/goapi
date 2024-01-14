// auth_handler.go
package handlers

import (
	"net/http"
	"os"

	"goapi/database"
	"goapi/models"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Get a login
// @Description Get a login
// @Tags login
// @Accept json
// @Produce json
// @Param loginInfo body models.Login true "Login Information"
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	var loginInfo models.Login
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
