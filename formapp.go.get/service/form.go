package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FormData struct {
	Name string `form:"name"`
}

func FormHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "form_name.html", nil)
}

func RegisterNameHandler(ctx *gin.Context) {
	var data FormData
	ctx.Bind(&data)
	ctx.HTML(http.StatusOK, "form_submit.html", &data)
}
