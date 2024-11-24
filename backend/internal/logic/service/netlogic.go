package service

import (
	"context"
	"net/http"
	"net/url"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/PuerkitoBio/goquery"
	"github.com/zeromicro/go-zero/core/logx"
)

type NetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取网站信息
func NewNetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NetLogic {
	return &NetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NetLogic) Net(req *types.IDRequest) (resp *types.Collection, err error) {
	resp, err = queryMeta(req.ID)
	if err != nil {
		return nil, err
	}

	return
}

func queryMeta(link string) (*types.Collection, error) {

	URL, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// 检查响应状态码
	if res.StatusCode != 200 {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	title := doc.Find("title").Text()
	description := doc.Find("meta[name='description']").AttrOr("content", "")

	return &types.Collection{
		CID:         URL.Host,
		Title:       title,
		Link:        link,
		Description: description,
	}, nil
}
