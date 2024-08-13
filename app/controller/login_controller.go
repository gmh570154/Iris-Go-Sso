package controller

import (
	"iris_master/common/configs"
	"iris_master/log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/requestid"
	"github.com/kataras/iris/v12/sessions"
)

func Login(ctx iris.Context) {
	log.Log.Info("user login")
	session := sessions.Get(ctx)
	session.Set("authenticated", true)  // todo 设置用户登录会话
	session.Set("session_id", "123456") // todo 设置用户登录会话

	request_id := requestid.Get(ctx)
	log.Log.Info(request_id)

	glb_request_id := ctx.GetHeader("glb_request_id")
	log.Log.Info(glb_request_id)
	callback_url := "redirect_uri=" + configs.AppConfig.Console.HomeUrl + "&response_type=" + configs.AppConfig.Sso.GrantType

	ctx.Redirect(configs.AppConfig.Sso.LoginUrl+"?"+callback_url, iris.StatusTemporaryRedirect)
}

func Logout(ctx iris.Context) {
	log.Log.Info("user logout")
	session := sessions.Get(ctx)
	// Revoke users authentication
	sid := session.ID()
	log.Log.Info("logout session id: " + sid)
	session.Set("authenticated", false) // todo 删除用户会话
	session.Destroy()

	callback_url := "redirect_uri=" + configs.AppConfig.Console.HomeUrl + "&response_type=" + configs.AppConfig.Sso.GrantType + "&client_id=" + configs.AppConfig.Sso.ClientId

	ctx.Redirect(configs.AppConfig.Sso.LogoutUrl+"?"+callback_url, iris.StatusTemporaryRedirect)

}
