// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package home

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 首页推荐商品
func NewHomeRecommendProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendProductsLogic {
	return &HomeRecommendProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendProductsLogic) HomeRecommendProducts(req *types.RecommendProductsReq) (resp *types.RecommendProductsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
