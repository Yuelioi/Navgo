package main

import (
	"context"
	"net/http"

	"backend/internal/biz"
	"backend/internal/global"
	"backend/internal/handler"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	_ "backend/internal/db"
)

func main() {

	server := rest.MustNewServer(global.ConfInst.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, global.SVCInst)

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

	server.Start()
}
