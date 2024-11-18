package collection

import (
	"context"
	"encoding/json"
	"os"

	"backend/internal/common/db"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 单页面
func NewCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectionLogic {
	return &CollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type CC struct {
	ID      string `json:"id"`
	Favicon string `json:"favicon"`
}

func (l *CollectionLogic) Collection(req *types.AnyRequest) (resp *types.Collection, err error) {
	// todo: add your logic here and delete this line

	var cos []*types.Collection
	db.DB.Model(types.Collection{}).Find(&cos)

	result := make([]*CC, 0)

	for _, c := range cos {
		result = append(result, &CC{c.CID, c.Favicon})
	}

	data, _ := json.Marshal(result)

	f, _ := os.Create("demo.json")
	defer f.Close()

	f.Write(data)
	return
}
