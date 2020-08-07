package handlers

import (
	"net/http"

	"github.com/elmawardy/tahor/models"
	"github.com/gin-gonic/gin"
)

func UserLogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := struct {
			JWT string
		}{}
		c.Bind(&requestData)

		err := models.UserLogout(requestData.JWT)
		if err != nil {
			c.JSON(http.StatusOK, map[string]string{
				"State": "Error",
				"Msg":   "Error Logging out",
			})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"State": "Success",
		})
	}
}

func UserGetBasicInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo := models.UserBasicInfo{}
		requestData := struct {
			JWT string
		}{}
		c.Bind(&requestData)
		userInfo, err := models.UserGetBasicInfo(requestData.JWT)
		if err != nil {
			c.JSON(http.StatusOK, map[string]string{
				"State": "Error",
				"Msg":   "Error Logging in, check username and password",
			})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"State":     "Success",
			"user_info": userInfo,
		})
	}
}

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
