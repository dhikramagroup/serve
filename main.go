package main

import (
	"time"
	"web/dhikrama/src/modules"
	"web/dhikrama/src/repository"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/cache"
)

func main() {
	app := modules.DhikramaApp()
	app.Use(cache.StaticCache(24 * time.Hour))

	app.RegisterView(iris.HTML("./web/views", ".html"))

	app.Get("/", repository.HomePage)
	app.Get("/services-us", repository.ServicesPage)
	app.Get("/contact-us", repository.ContactPage)
	app.Get("/about-us", repository.AboutPage)
	// app.Listen(":8080")
	app.Run(iris.AutoTLS(":443", "kang-bangunan.com www.kang-bangunan.com", "dhirama.group@gmail.com"))
}
