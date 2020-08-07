package main

import (
	"time"

	"github.com/elmawardy/tahor/global"
	"github.com/elmawardy/tahor/handlers"
	"github.com/elmawardy/tahor/models"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/jinzhu/gorm"
)

func main() {

	// Check Database and migrations

	var err error
	global.DB, err = gorm.Open(global.Config["driver"], global.Config["constring"])
	defer global.DB.Close()
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	global.DB.AutoMigrate(&models.Case{}, &models.Organization{}, &models.Tag{}, &models.User{})

	//  Web Server

	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	r.POST("/api/case/insert", handlers.CaseInsertHandler())
	r.GET("/api/case/select", handlers.CaseSelectHandler())

	r.POST("/api/user/register", handlers.UserRegisterHandler())
	r.POST("/api/user/login", handlers.UserLoginHandler())
	r.POST("/api/user/getbasicinfo", handlers.UserGetBasicInfoHandler())
	r.POST("/api/user/logout", handlers.UserLogoutHandler())

	r.Run(":8000")
}
