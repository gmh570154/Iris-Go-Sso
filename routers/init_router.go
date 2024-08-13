package routers

import (
	"iris_master/app/controller"

	"iris_master/middlerware"

	"github.com/iris-contrib/middleware/csrf"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/requestid"
)

func InitRouter(app *iris.Application) {
	CSRF := csrf.Protect(
		[]byte("9AB0F421E53A477C084477AEA06096F5"),
		csrf.Secure(false), // WARNING: Set it to true on production with HTTPS.
	)

	userLoginout := app.Party("/", middlerware.CommonMiddlerWare)
	userLoginout.Use(requestid.New()) // 设置header  x_request_id
	//userAPI.Use(CSRF)   // 启用资源跨域限制
	userLoginout.Get("login", controller.Login)
	userLoginout.Get("logout", controller.Logout)

	userAPI := app.Party("/api", middlerware.LoginAuthorize) //登录鉴权
	//userAPI.Use(CSRF)
	userAPI.Use(requestid.New())
	userAPI.Get("/test", controller.TestController, middlerware.CommonMiddlerWare)

	userREST := app.Party("/rest") //登录鉴权  //, middlerware.LoginAuthorize
	//userAPI.Use(CSRF)
	userREST.Use(requestid.New())
	userREST.Get("/user", controller.UserList, middlerware.CommonMiddlerWare)
	userREST.Get("/keepalive", controller.KeepAlive, middlerware.CommonMiddlerWare)
	userREST.Get("/logout", controller.Logout)
	userREST.Get("/uservar", controller.UserVar)
	userREST.Get("/videos", controller.Videos)
	userREST.Get("/video", controller.Video)

	// csrf 测试
	userAPI2 := app.Party("/user")
	//userAPI2.Use(CSRF)
	userAPI2.Use(requestid.New())
	userAPI2.Get("/signup", controller.GetSignupForm)
	// POST requests without a valid token will return a HTTP 403 Forbidden.
	userAPI2.Post("/signup", controller.PostSignupForm)
	// Remove the CSRF middleware (1)
	userAPI2.Post("/unprotected", controller.Unprotected).RemoveHandler(CSRF) // or RemoveHandler("iris-contrib.csrf.token")

}
