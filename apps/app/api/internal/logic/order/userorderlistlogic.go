// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package order

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户订单列表
func NewUserOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOrderListLogic {
	return &UserOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserOrderListLogic) UserOrderList(req *types.UserOrderListReq) (resp *types.UserOrderListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
