package middlerware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"iris_master/common/models"
	"iris_master/log"
	"net/http"
	"path"
	"time"

	"github.com/kataras/iris/v12"
)

func CommonMiddlerWare(ctx iris.Context) {
	models.GeneratorGlbRequestId(ctx)
	p := ctx.Request().URL.Path
	method := ctx.Request().Method
	start := time.Now()
	fields := make(map[string]interface{})
	fields["title"] = "请求日志"
	fields["fun_name"] = path.Join(method, p)
	fields["ip"] = ctx.Request().RemoteAddr
	fields["method"] = method
	fields["url"] = ctx.Request().URL.String()
	fields["proto"] = ctx.Request().Proto
	//fields["header"] = ctx.Request().Header
	fields["user_agent"] = ctx.Request().UserAgent()
	fields["x_request_id"] = ctx.GetHeader("X-Request-Id")
	fields["glb_request_id"] = ctx.GetHeader("glb_request_id") // 根据需要透传到底层微服务参数

	// 如果是POST/PUT请求，并且内容类型为JSON，则读取内容体
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err == nil {
			defer ctx.Request().Body.Close()
			buf := bytes.NewBuffer(body)
			ctx.Request().Body = ioutil.NopCloser(buf)
			fields["content_length"] = ctx.GetContentLength()
			fields["body"] = string(body)
		}
	}
	ctx.Next()

	//下面是返回日志
	fields["res_status"] = ctx.ResponseWriter().StatusCode()
	timeConsuming := time.Since(start).Nanoseconds() / 1e6
	msg := fmt.Sprintf("[http] %s-%s-%s-%d(%dms)",
		p, ctx.Request().Method, ctx.Request().RemoteAddr, ctx.ResponseWriter().StatusCode(), timeConsuming)
	log.Log.Debug(fields)
	log.Log.Infof(msg)
	ctx.Next() // execute the next handler, in this case the main one.
}

func Before(ctx iris.Context) {
	requestPath := ctx.Path()
	log.Log.Info("Before the mainHandler: " + requestPath)
	ctx.Next() // execute the next handler, in this case the main one.
}

func After(ctx iris.Context) {
	log.Log.Info("After the mainHandler")
	ctx.Next()
}

func MainHandler(ctx iris.Context) {
	log.Log.Info("Inside mainHandler")
	ctx.Next() // execute the "after".
}
