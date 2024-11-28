package main

import (
	"context"
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"

	"backend/internal/common/biz"
	"backend/internal/common/constants"
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	_ "backend/internal/common/db"
)

const basename = "/web" // 虚拟路由根路径

//go:embed public
var assets embed.FS

type NotFoundHandler struct {
	fs         http.FileSystem
	fileServer http.Handler
}

func (n NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := strings.TrimPrefix(path.Clean(r.URL.Path), basename)
	if len(filePath) == 0 {
		filePath = basename
	}

	file, err := n.fs.Open(filePath)
	switch {
	case err == nil:
		n.fileServer.ServeHTTP(w, r)
		_ = file.Close()
		return
	case os.IsNotExist(err):
		r.URL.Path = "/"
		n.fileServer.ServeHTTP(w, r)
		return
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
}

func main() {
	logc.Info(context.Background(), "启动服务中")

	sub, _ := fs.Sub(assets, "public")
	fs := http.FS(sub)
	fileServer := http.FileServer(fs)

	server := rest.MustNewServer(

		constants.ConfInst.RestConf,
		rest.WithNotFoundHandler(&NotFoundHandler{ // 自定义 NotFoundHandler，对虚拟路由做处理
			fs:         fs,
			fileServer: fileServer,
		}),
	)

	// 跨域中间件
	server.Use(middleware.CORS())
	server.Use(middleware.Limit())

	defer server.Stop()

	handler.RegisterHandlers(server, constants.SVCInst)

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
