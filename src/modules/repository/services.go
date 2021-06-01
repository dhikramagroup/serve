package repository

import (
	"github.com/kataras/iris/v12"
)

func ServicesPage(ctx iris.Context) {
	data := iris.Map{
		"Title":      "Page Title",
		"FooterText": "Footer contents",
		"Message":    "Main contents",
	}

	ctx.ViewLayout("layouts/main")
	ctx.View("services", data)
}
