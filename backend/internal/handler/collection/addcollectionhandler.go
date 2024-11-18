package collection

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"backend/internal/common/cache"
	global "backend/internal/common/constants"
	"backend/internal/logic/collection"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func saveFile(w http.ResponseWriter, r *http.Request, file multipart.File, dir, cid string) (string, error) {
	// 保存文件到本地

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 如果目录不存在，创建目录
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", err
		}

	}

	target := filepath.Join(dir, cid)
	outFile, err := os.Create(target + ".png")
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, errors.New("unable to create file"))
		return "", err

	}
	defer outFile.Close()

	// 将文件内容保存到本地
	_, err = io.Copy(outFile, file)
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, errors.New("unable to save file"))
		return "", err
	}
	return "", nil
}

// 解析表单
func parseForm(r *http.Request) (*types.Collection, error) {
	link := r.FormValue("link")
	title := r.FormValue("title")
	proxy, _ := strconv.ParseBool(r.FormValue("proxy"))
	description := r.FormValue("description")

	linkRaw, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	// 获取当前分类的页面数
	order := 1

	// 获取分类
	cat := &types.Category{}

	return &types.Collection{
		CID:         linkRaw.Host,
		Title:       title,
		Link:        link,
		Order:       order,
		Path:        "",
		Proxy:       proxy,
		CategoryID:  "",
		Category:    cat,
		Description: description,
		Favicon:     "",
		Tags:        []string{},
		View:        0,
	}, nil
}

// 增加页面
func AddCollectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req *types.Collection

		// 解析 multipart 表单
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("无法解析字段"))
			return
		}

		// 获取表单字段
		req, err := parseForm(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("无法解析链接"))
			return
		}

		// 验证是否存在

		if cache.Cache.Exists(req.CID) {
			httpx.ErrorCtx(r.Context(), w, errors.New("已经提交过啦,正在审核中"))
			return
		}

		// 保存文件
		file, _, err := r.FormFile("favicon")
		if err == nil {
			defer file.Close()
			icon, err := saveFile(w, r, file, global.ConfInst.Resource.Icons, req.CID)
			if err != nil {
				return
			}
			req.Favicon = icon
		}

		l := collection.NewAddCollectionLogic(r.Context(), svcCtx)
		resp, err := l.AddCollection(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
