package main

import (
	"time"
	"web/dhikrama/src/modules/repository"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.HandleDir("/", iris.Dir("./web/public"))

	app.RegisterView(iris.HTML("./web/views", ".html"))

	app.Get("/", iris.StaticCache(24*time.Hour), repository.HomePage)
	app.Get("/services-us", repository.ServicesPage)
	app.Get("/contact-us", repository.ContactPage)
	app.Get("/about-us", repository.AboutPage)
	app.Run(iris.AutoTLS(":443", "kang-bangunan.com www.kang-bangunan.com", "dhirama.group@gmail.com"))
}
