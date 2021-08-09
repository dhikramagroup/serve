package modules

import (
	"regexp"

	"github.com/kataras/iris/v12"
)

var dirOptions = iris.DirOptions{
	IndexName: "index.html",
	PushTargetsRegexp: map[string]*regexp.Regexp{
		"/":       iris.MatchCommonAssets,
		"/js":     iris.MatchCommonAssets,
		"/css":    iris.MatchCommonAssets,
		"/images": iris.MatchCommonAssets},
	// The `Compress` field is ignored
	// when the file is cached (when Cache.Enable is true),
	// because the cache file has a map of pre-compressed contents for each encoding
	// that is served based on client's accept-encoding.
	Compress: true, // true or false does not matter here.
	ShowList: true,
	SPA:      true,
	Cache: iris.DirCacheOptions{
		Enable:         true,
		CompressIgnore: iris.MatchImagesAssets,
		// Here, define the encodings that the cached files should be pre-compressed
		// and served based on client's needs.
		Encodings:       []string{"gzip", "deflate", "br", "snappy"},
		CompressMinSize: 50, // files smaller than this size will NOT be compressed.
		Verbose:         1,
	},
}

func DhikramaApp() *iris.Application {
	app := iris.New()

	app.Favicon("./web/public/favicon.ico")
	// app.Use(cache.StaticCache(24 * time.Hour))

	// first parameter is the request path
	// second is the system directory
	//
	// app.HandleDir("/css", iris.Dir("./assets/css"))
	// app.HandleDir("/js",  iris.Dir("./assets/js"))
	app.HandleDir("/js", iris.Dir("./web/public/js"), dirOptions)
	app.HandleDir("/css", iris.Dir("./web/public/css"), dirOptions)
	// app.HandleDir("/icons", iris.Dir("./web/public/icons"), dirOptions)
	app.HandleDir("/images", iris.Dir("./web/public/images"), dirOptions)
	app.HandleDir("/", iris.Dir("./web/public/robots"))

	// You can also register any index handler manually, order of registration does not matter:
	// v1.Get("/static", [...custom middleware...], func(ctx iris.Context) {
	//  [...custom code...]
	// 	ctx.ServeFile("./assets/index.html")
	// })

	// http://localhost:8080/v1/static
	// http://localhost:8080/v1/static/css/main.css
	// http://localhost:8080/v1/static/js/jquery-2.1.1.js
	// http://localhost:8080/v1/static/favicon.ico
	return app
}
