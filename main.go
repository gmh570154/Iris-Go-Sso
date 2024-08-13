package main

import (
	"iris_master/common/configs"
	"iris_master/log"
	"iris_master/middlerware"
	"iris_master/routers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

const cookieNameForSessionID = "session_id_cookie"

func main() {
	app := iris.New()

	configs.InitConfig()

	logConfig := configs.LogConfig{
		Level: configs.AppConfig.LogLevel,
		Path:  "./log", //日志文件夹路径
		Save:  5,       // 日志备份数
	}
	log.InitLogger(logConfig)

	app.Logger().SetLevel("debug")

	htmlEngine := iris.HTML("./views", ".html")
	app.RegisterView(htmlEngine)

	sess := sessions.New(sessions.Config{Cookie: cookieNameForSessionID, AllowReclaim: true})
	app.Use(sess.Handler())

	routers.InitRouter(app)

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}
		ctx.Writef("(Unexpected) internal server error")
	})

	app.Get("/", middlerware.Before, middlerware.MainHandler, middlerware.After)

	app.Use(func(ctx iris.Context) {
		log.Log.Info(`before the party's routes and its children`)
		ctx.Next()
	})

	app.Done(func(ctx iris.Context) {
		log.Log.Info("this is executed always last")
		message := ctx.Values().GetString("message")
		log.Log.Info("message: " + message)
	})

	cfg := iris.YAML("./configs/app.yml")
	host := cfg.Other["Host"].(string) + cfg.Other["Port"].(string)

	app.Listen(host)
}
