// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package cart

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除购物车商品
func NewDeleteCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCartLogic {
	return &DeleteCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCartLogic) DeleteCart(req *types.DeleteCartReq) (resp *types.DeleteCartResp, err error) {
	// todo: add your logic here and delete this line

	return
}
