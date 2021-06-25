package repository

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func AboutPage(ctx iris.Context) {

	data := iris.Map{
		"canonical":  "https://www.kang-bangunan.com/about-us",
		"Title":      "Page Title",
		"FooterText": "Footer contents",
		"Message":    "Main contents",
	}

	target := "/about.html"

	if pusher, ok := ctx.ResponseWriter().Naive().(http.Pusher); ok {
		err := pusher.Push(target, nil)
		if err != nil {
			if err == iris.ErrPushNotSupported {
				ctx.StopWithText(iris.StatusHTTPVersionNotSupported, "HTTP/2 push not supported.")
			} else {
				ctx.StopWithError(iris.StatusInternalServerError, err)
			}
			return
		}

		ctx.ViewLayout("layouts/main")
		ctx.View("about", data)
	}
}
