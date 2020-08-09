package handlers

import (
	"fmt"
	"net/http"

	"github.com/elmawardy/tahor/models"
	"github.com/gin-gonic/gin"
)

func VoteUpHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		requestData := struct {
			CaseID int
			JWT    string
		}{}
		c.Bind(&requestData)

		vote := models.Vote{}

		user := models.User{}

		claims, err := models.ParseToken(requestData.JWT)
		if claims == nil {
			// log.Println(err)
			fmt.Println(`{"Status":"Error","Msg":"Unauthorized"}`)
		}
		email := fmt.Sprintf("%v", claims["email"])
		user.Email = email

		err = user.InjectIdByJwt(requestData.JWT)

		if err != nil {
			c.JSON(http.StatusOK, map[string]string{
				"State": "Error",
				"Msg":   "Error Getting user id",
			})
			return
		}

		err = vote.VoteUp(user.ID, requestData.CaseID)

		if err != nil {
			c.JSON(http.StatusOK, map[string]string{
				"State": "Error",
				"Msg":   "Error Voting",
			})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"State": "Success",
		})

	}
}
