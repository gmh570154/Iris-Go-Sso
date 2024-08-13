package controller

import (
	"iris_master/common/configs"
	"iris_master/common/models"
	"iris_master/log"

	"github.com/iris-contrib/middleware/csrf"
	"github.com/kataras/iris/v12"
)

func UserList(ctx iris.Context) {
	log.Log.Info("list users")
	// take the info from the "before" handler.
	info := "list users"

	// write something to the client as a response.
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // execute the "after".
}

func GetSignupForm(ctx iris.Context) {
	ctx.ViewData(csrf.TemplateTag, csrf.TemplateField(ctx))
	ctx.View("user/signup.html")
}

func PostSignupForm(ctx iris.Context) {
	ctx.Writef("You're welcome mate!")
}

func Unprotected(ctx iris.Context) {
	ctx.Writef("Hey, I am open to CSRF attacks!")
}

func KeepAlive(ctx iris.Context) {
	log.Log.Info("keep alive")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With, Current-Page")
	ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	//return &models.Response{1, "success"}
	ctx.JSON(iris.Map{"message": "hello word", "code": 200})
}

func Video(ctx iris.Context) {
	log.Log.Info("get video url")

	video_name := ctx.URLParam("video_name")
	log.Log.Info("params: " + video_name)

	video_url := ""
	for i := range configs.AppConfig.Videos {
		if configs.AppConfig.Videos[i].Name == video_name {
			video_url = configs.AppConfig.Videos[i].Url
		}
	}

	data := iris.Map{
		"video_url": video_url,
	}
	response := models.GenSuccessData(data)
	ctx.JSON(response)

}

func Videos(ctx iris.Context) {
	log.Log.Info("list video url")

	url_list := make([]interface{}, len(configs.AppConfig.Videos)) //初始化切片

	for i := range configs.AppConfig.Videos {
		url_list[i] = iris.Map{
			"video_url": configs.AppConfig.Videos[i].Url,
		}
	}

	response := models.GenSuccessData(url_list)
	ctx.JSON(response)

}

func UserVar(ctx iris.Context) {
	log.Log.Info("user var api")

	uservar := models.UserVar{
		HomeUrl:   configs.AppConfig.Console.HomeUrl,
		LoginUrl:  configs.AppConfig.Sso.LoginUrl,
		LogoutUrl: configs.AppConfig.Sso.LogoutUrl,
		GrantType: configs.AppConfig.Sso.GrantType,
	}

	response := models.GenSuccessData(uservar)
	ctx.JSON(response)
}
