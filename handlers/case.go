package handlers

import (
	"net/http"

	"github.com/elmawardy/tahor/models"
	"github.com/gin-gonic/gin"
)

func CaseInsertHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		icase := models.Case{}
		c.Bind(&icase)
		err := icase.Insert()

		if err != nil {
			c.JSON(http.StatusOK, map[string]string{
				"State": "Error",
				"Msg":   "Error inserting in DB",
			})
		}

		c.JSON(http.StatusOK, map[string]string{
			"State": "Fail",
		})
	}
}

func CaseSelectHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		cases, err := models.SelectCases(0, 6)

		if err != nil {
			c.JSON(http.StatusOK, map[string]string{
				"State": "Error",
				"Msg":   "Error inserting in DB",
			})
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"State": "Success",
			"Cases": cases,
		})
	}
}
