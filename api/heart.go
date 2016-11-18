package api

import (
	"../model"
	"../service"

	"fmt"

	"github.com/kataras/iris"
)

type HeartAPI struct {
	*iris.Context
}

var (
	heart = new(service.HeartService)
)

// List GET /heart/:duration
func (api HeartAPI) List(ctx *iris.Context) {

	duration, _ := ctx.ParamInt("duration")

	result, err := heart.List(duration)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(iris.StatusOK, model.RsMsg("1"))
	} else {
		ctx.JSON(iris.StatusOK, result)
	}
}

func (api HeartAPI) Delete(ctx *iris.Context) {
	heartData := model.HeartData{}
	ctx.ReadJSON(&heartData)
	result, err := heart.Delete(heartData)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(iris.StatusOK, model.RsMsg("1"))
	} else {
		ctx.JSON(iris.StatusOK, result)
	}
}

func (api HeartAPI) Save(ctx *iris.Context) {
	heartData := model.HeartData{}
	ctx.ReadJSON(&heartData)
	result, err := heart.Save(heartData)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(iris.StatusOK, model.RsMsg("1"))
	} else {
		ctx.JSON(iris.StatusOK, result)
	}
}
