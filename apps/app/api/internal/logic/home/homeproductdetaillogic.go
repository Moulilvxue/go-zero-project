// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package home

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 商品详情
func NewHomeProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeProductDetailLogic {
	return &HomeProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeProductDetailLogic) HomeProductDetail(req *types.ProductDetailReq) (resp *types.ProductDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
