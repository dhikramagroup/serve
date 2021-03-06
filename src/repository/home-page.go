package repository

import (
	"github.com/kataras/iris/v12"
)

func HomePage(ctx iris.Context) {

	data := iris.Map{
		"canonical":  "https://www.kang-bangunan.com",
		"Title":      "Home page || dikrama jasa",
		"FooterText": "Footer contents",
		"Message":    "Main contents",
	}

	ctx.ViewLayout("layouts/main")
	ctx.View("index", data)

}
