// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package payment

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitiatePaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发起支付
func NewInitiatePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitiatePaymentLogic {
	return &InitiatePaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitiatePaymentLogic) InitiatePayment(req *types.InitiatePaymentReq) (resp *types.InitiatePaymentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
