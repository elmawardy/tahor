package handlers

import (
	"net/http"

	"github.com/elmawardy/tahor/models"
	"github.com/gin-gonic/gin"
)

func UserLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{}
		c.Bind(&user)

		jwtoken, err := user.Login()
		if err != nil {
			c.JSON(http.StatusOK, map[string]string{
				"State": "Error",
				"Msg":   "Error Logging in, check username and password",
			})
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"State": "Success",
			"jwt":   jwtoken,
		})
	}
}

func UserRegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := &models.User{}
		c.Bind(&user)
		err := user.Register()
		if err != nil {
			c.JSON(http.StatusOK, map[string]string{
				"State": "Error",
				"Msg":   "Error inserting in DB",
			})
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"State": "Success",
		})
	}
}
