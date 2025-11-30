// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package home

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeCategoryProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分类商品列表
func NewHomeCategoryProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeCategoryProductsLogic {
	return &HomeCategoryProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeCategoryProductsLogic) HomeCategoryProducts(req *types.CategoryProductsReq) (resp *types.CategoryProductsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
