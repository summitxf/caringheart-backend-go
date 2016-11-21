package api

import (
	"../model"

	"fmt"

	"github.com/kataras/iris"
)

type AuthAPI struct {
	*iris.Context
}

var (
// auth = new(service.AuthService)
)

func (api AuthAPI) Login(ctx *iris.Context) {
	authData := model.AuthData{}
	ctx.ReadJSON(&authData)
	fmt.Println(authData)
	authToken := "abcdef123456"

	ctx.JSON(iris.StatusOK, map[string]string{"token": authToken})

}
