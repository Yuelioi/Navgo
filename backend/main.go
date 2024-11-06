package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"backend/internal/biz"
	"backend/internal/config"
	"backend/internal/handler"
	"backend/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/main.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 统一错误处理
	httpx.SetErrorHandler(func(err error) (int, any) {
		// var e *biz.Error
		// switch {
		// case errors.As(err, &e):
		// 	return http.StatusOK, biz.Fail(e)
		// default:
		// 	return http.StatusInternalServerError, biz.NewError(500, e.Error())
		// }

		return http.StatusOK, biz.Result{
			Code: -1,
			Msg:  err.Error(),
			Data: nil,
		}
	})

	// 统一成功处理
	httpx.SetOkHandler(func(ctx context.Context, data any) any {
		return biz.Success(data)
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
