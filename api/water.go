package api

import (
	"../model"
	"../service"

	"fmt"

	"github.com/kataras/iris"
)

type WaterAPI struct {
	*iris.Context
}

var (
	water = new(service.WaterService)
)

// List GET /water/:duration
func (api WaterAPI) List(ctx *iris.Context) {

	duration, _ := ctx.ParamInt("duration")

	result, err := water.List(duration)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(iris.StatusOK, model.RsMsg("1"))
	} else {
		ctx.JSON(iris.StatusOK, result)
	}
}

func (api WaterAPI) Delete(ctx *iris.Context) {
	waterData := model.WaterData{}
	ctx.ReadJSON(&waterData)
	result, err := water.Delete(waterData)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(iris.StatusOK, model.RsMsg("1"))
	} else {
		ctx.JSON(iris.StatusOK, result)
	}
}

func (api WaterAPI) Save(ctx *iris.Context) {
	waterData := model.WaterData{}
	ctx.ReadJSON(&waterData)
	result, err := water.Save(waterData)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(iris.StatusOK, model.RsMsg("1"))
	} else {
		ctx.JSON(iris.StatusOK, result)
	}
}
