// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package cart

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddToCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加商品到购物车
func NewAddToCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddToCartLogic {
	return &AddToCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddToCartLogic) AddToCart(req *types.AddToCartReq) (resp *types.AddToCartResp, err error) {
	// todo: add your logic here and delete this line

	return
}
