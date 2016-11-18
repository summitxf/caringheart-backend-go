package main

import (
	"./api"

	"github.com/kataras/iris"
)

func main() {
	// set the custom errors
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("errors/404.html", iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("errors/500.html", nil, iris.RenderOptions{"layout": iris.NoLayout})
	})

	registerAPI()

	// start the server
	iris.Listen("127.0.0.1:8080")
}

func registerAPI() {

	caringheart := iris.Party("/caringheart-backend")

	heart := new(api.HeartAPI)
	caringheart.Get("/heart/:duration", heart.List)
	caringheart.Delete("/heart", heart.Delete)
	caringheart.Post("/heart", heart.Save)

}
