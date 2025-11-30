package logic

import (
	"context"
	"product/rpc/product"

	"product/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAndUpdateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAndUpdateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAndUpdateStockLogic {
	return &CheckAndUpdateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckAndUpdateStockLogic) CheckAndUpdateStock(in *product.CheckAndUpdateStockRequest) (*product.CheckAndUpdateStockResponse, error) {
	// todo: add your logic here and delete this line

	return &product.CheckAndUpdateStockResponse{}, nil
}
