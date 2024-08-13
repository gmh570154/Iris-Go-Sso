package controller

import (
	"iris_master/log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func TestController(ctx iris.Context) {
	log.Log.Info("Inside controller")
	session := sessions.Get(ctx)

	session_id := session.GetString("session_id")
	log.Log.Info("session_id: " + session_id)
	// write something to the client as a response.
	ctx.HTML("<h1>Response</h1>")
	ctx.Next() // execute the "after".
}
