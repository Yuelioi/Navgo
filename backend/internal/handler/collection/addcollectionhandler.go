package collection

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"backend/internal/common/biz"
	"backend/internal/common/cache"
	global "backend/internal/common/constants"
	"backend/internal/common/dao/user"
	"backend/internal/logic/collection"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func saveFile(file multipart.File, dir, cid string) error {
	// 保存文件到本地

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 如果目录不存在，创建目录
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}

	}

	target := filepath.Join(dir, cid)
	outFile, err := os.Create(target + ".png")
	if err != nil {

		return errors.New("unable to create file")

	}
	defer outFile.Close()

	// 将文件内容保存到本地
	_, err = io.Copy(outFile, file)
	if err != nil {
		return errors.New("unable to save file")
	}
	return nil
}

// 增加页面
func AddCollectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CollectionUpdateParams

		// 获取表单字段
		err := httpx.Parse(r, &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("无法解析链接"))
			return
		}

		isAdmin := false
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			username, err := biz.Validate(authorizationHeader)
			if err == nil {
				u, err := user.Get(username)
				if err == nil {
					if u.Role == "admin" {
						isAdmin = true
					}
				}
			}
		}

		u, err := url.Parse(req.Link)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("无法解析链接"))
			return
		}

		req.CID = u.Host

		if !isAdmin {
			// 验证审核区是否存在
			controller := cache.Manager.GetController(cache.InReviewCacheID)
			if controller != nil && controller.Exists(req.CID) {
				httpx.ErrorCtx(r.Context(), w, errors.New("已经提交过啦,正在审核中"))
				return
			}
		}

		// 保存文件
		file, _, err := r.FormFile("favicon")
		if err == nil {
			defer file.Close()
			err := saveFile(file, global.ConfInst.Resource.Icons, req.CID)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
		}

		l := collection.NewAddCollectionLogic(r.Context(), svcCtx)
		resp, err := l.AddCollection(&req, isAdmin)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
