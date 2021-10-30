package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FormData struct {
	Name string `form:"name"`
	Date  string `form:"date"`
	Message string `form:"message"`
}

func FormHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "form_start.html", nil)
}

func RegisterStartHandler(ctx *gin.Context) {
	var data FormData
	ctx.Bind(&data)
	ctx.HTML(http.StatusOK, "form_name.html", &data)
}

func RegisterNameHandler(ctx *gin.Context) {
	var data FormData
	ctx.Bind(&data)
	ctx.HTML(http.StatusOK, "form_date.html", &data)
}

func RegisterDateHandler(ctx *gin.Context) {
	var data FormData
	ctx.Bind(&data)
	ctx.HTML(http.StatusOK, "form_message.html", &data)
}

func RegisterMessageHandler(ctx *gin.Context) {
	var data FormData
	ctx.Bind(&data)
	ctx.HTML(http.StatusOK, "form_submit.html", &data)
}

func BackNameHandler(ctx *gin.Context) {
	var data FormData
	ctx.Bind(&data)
	ctx.HTML(http.StatusOK, "form_start.html", &data)
}

func BackDateHandler(ctx *gin.Context) {
	var data FormData
	ctx.Bind(&data)
	ctx.HTML(http.StatusOK, "form_name.html", &data)
}

func BackMessageHandler(ctx *gin.Context) {
	var data FormData
	ctx.Bind(&data)
	ctx.HTML(http.StatusOK, "form_date.html", &data)
}

func BackSubmitHandler(ctx *gin.Context) {
	var data FormData
	ctx.Bind(&data)
	ctx.HTML(http.StatusOK, "form_message.html", &data)
}
