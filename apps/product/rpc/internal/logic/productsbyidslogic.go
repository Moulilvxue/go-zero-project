package logic

import (
	"context"

	"product/rpc/internal/svc"
	"product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductsByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsByIdsLogic {
	return &ProductsByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsByIdsLogic) ProductsByIds(in *product.ProductsRequest) (*product.ProductsResponse, error) {
	// todo: add your logic here and delete this line

	return &product.ProductsResponse{}, nil
}
