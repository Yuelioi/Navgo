package collection

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"backend/internal/logic/collection"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func saveFile(w http.ResponseWriter, file multipart.File, cid string) (string, error) {
	// 保存文件到本地

	target := filepath.Join(cid)
	outFile, err := os.Create(target + ".png")
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return "", err

	}
	defer outFile.Close()

	// 将文件内容保存到本地
	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return "", err
	}
	return "", nil
}

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
			http.Error(w, "Unable to parse multipart form", http.StatusBadRequest)
			return
		}

		// 获取表单字段
		req, err := parseForm(r)
		if err != nil {
			http.Error(w, "Unable to parse link", http.StatusInternalServerError)
			return
		}
		file, _, err := r.FormFile("favicon")
		if err == nil {
			defer file.Close()
			icon, err := saveFile(w, file, req.CID)
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
