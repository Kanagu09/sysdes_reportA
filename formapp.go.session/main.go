package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"formapp.go/service"
)

// config
const port = 8000

func main() {
	// initialize Gin engine
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*.html")

	// routing
	engine.GET("/", rootHandler)
	engine.GET("/form", service.FormHandler)
	engine.POST("/form", service.FormHandler)
	engine.POST("/register-start", service.RegisterStartHandler)
	engine.POST("/register-name", service.RegisterNameHandler)
	engine.POST("/register-date", service.RegisterDateHandler)
	engine.POST("/register-message", service.RegisterMessageHandler)
	engine.POST("/back-name", service.BackNameHandler)
	engine.POST("/back-date", service.BackDateHandler)
	engine.POST("/back-message", service.BackMessageHandler)
	engine.POST("/back-submit", service.BackSubmitHandler)

	// start server
	engine.Run(fmt.Sprintf(":%d", port))
}

func rootHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "form_start.html", nil)
}
