package middlerware

import (
	"iris_master/common/configs"
	"iris_master/utils"

	"iris_master/log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func LoginAuthorize(ctx iris.Context) {

	log.Log.Info("start login author")
	session := sessions.Get(ctx)

	if auth, _ := session.GetBoolean("authenticated"); auth {
		if sid := session.ID(); sid != "" {
			log.Log.Info("user login, session id: " + sid)
		}
	} else {
		if code := ctx.Request().URL.Query().Get("code"); code != "" { // code验证登录
			redirect_uri := "http://" + ctx.Request().Host + ctx.Path()
			result := utils.CodeOauth(code, redirect_uri)
			if result {
				log.Log.Info("set cdoe session true")
				session.Set("authenticated", true) // todo 设置用户登录会话
			} else {
				log.Log.Info("set code session false")
				session.Set("authenticated", false)
				ctx.StatusCode(iris.StatusForbidden)
				return
			}

		} else if token := ctx.Request().URL.Query().Get("access_token"); token != "" { // token 验证登录
			redirect_uri := "http://" + ctx.Request().Host + ctx.Path()
			result := utils.TokenOauth(token, redirect_uri)
			if result {
				log.Log.Info("set token session true")
				session.Set("authenticated", true) // todo 设置用户登录会话
			} else {
				log.Log.Info("set token session false")
				session.Set("authenticated", true)
				ctx.StatusCode(iris.StatusForbidden)
				return
			}
		} else {
			// 未认证，重定向到登录页面
			log.Log.Info("query have not code or token")
			ctx.StatusCode(iris.StatusUnauthorized)
			sid := session.ID()
			log.Log.Info("session id: " + sid)
			session.Set("authenticated", false)
			ctx.JSON(iris.Map{"sso_url": configs.AppConfig.Sso.LoginUrl, "code": 401})
			return
		}
		// todo 用户名密码登录
	}
	log.Log.Info("end login author")

	ctx.Next() // execute the next handler, in this case the main one.
}
