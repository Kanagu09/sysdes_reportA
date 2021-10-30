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
	engine.GET("/register-name", service.RegisterNameHandler)

	// start server
	engine.Run(fmt.Sprintf(":%d", port))
}

func rootHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "form_start.html", nil)
}
