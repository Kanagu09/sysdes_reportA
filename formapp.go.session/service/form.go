package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FormData struct {
	Name    string `form:"name"`
	Date    string `form:"date"`
	Message string `form:"message"`
}

type Data struct {
	Name    string ""
	Date    string ""
	Message string ""
}

var form_data = make(map[string]*Data)

func checkCookieId(ctx *gin.Context) string {
	id, _ := ctx.Request.Cookie("id")
	return id.Value
}

func setCookie(ctx *gin.Context, uuid string) {
	_, err := ctx.Request.Cookie("id")
	if err != nil {
	} else {
		ctx.SetCookie("id", "", -1, "/", "localhost", false, true)
		fmt.Print("Disable Cookie.", "\n")
	}
	ctx.SetCookie("id", uuid, 100, "/", "localhost", false, true)
	fmt.Print("Set Cookie: ", uuid, "\n")
}

func FormHandler(ctx *gin.Context) {
	// generate uuid
	uuid_val, _ := uuid.NewRandom()
	uuid := uuid_val.String()
	// set cookie
	setCookie(ctx, uuid)
	// make data variable
	form_data[uuid] = &Data{}
	ctx.HTML(http.StatusOK, "form_start.html", nil)
}

func RegisterStartHandler(ctx *gin.Context) {
	uuid := checkCookieId(ctx)
	ctx.HTML(http.StatusOK, "form_name.html", form_data[uuid])
}

func RegisterNameHandler(ctx *gin.Context) {
	var data FormData
	ctx.Bind(&data)
	uuid := checkCookieId(ctx)
	form_data[uuid].Name = data.Name
	ctx.HTML(http.StatusOK, "form_date.html", form_data[uuid])
}

func RegisterDateHandler(ctx *gin.Context) {
	var data FormData
	ctx.Bind(&data)
	uuid := checkCookieId(ctx)
	form_data[uuid].Date = data.Date
	ctx.HTML(http.StatusOK, "form_message.html", form_data[uuid])
}

func RegisterMessageHandler(ctx *gin.Context) {
	var data FormData
	ctx.Bind(&data)
	uuid := checkCookieId(ctx)
	form_data[uuid].Message = data.Message
	ctx.HTML(http.StatusOK, "form_submit.html", form_data[uuid])
}

func BackNameHandler(ctx *gin.Context) {
	uuid := checkCookieId(ctx)
	ctx.HTML(http.StatusOK, "form_start.html", form_data[uuid])
}

func BackDateHandler(ctx *gin.Context) {
	uuid := checkCookieId(ctx)
	ctx.HTML(http.StatusOK, "form_name.html", form_data[uuid])
}

func BackMessageHandler(ctx *gin.Context) {
	uuid := checkCookieId(ctx)
	ctx.HTML(http.StatusOK, "form_date.html", form_data[uuid])
}

func BackSubmitHandler(ctx *gin.Context) {
	uuid := checkCookieId(ctx)
	ctx.HTML(http.StatusOK, "form_message.html", form_data[uuid])
}
