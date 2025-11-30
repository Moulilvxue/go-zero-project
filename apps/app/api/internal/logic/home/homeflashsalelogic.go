// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package home

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeFlashSaleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 首页抢购
func NewHomeFlashSaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeFlashSaleLogic {
	return &HomeFlashSaleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeFlashSaleLogic) HomeFlashSale() (resp *types.FlashSaleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
