package logic

import (
	"context"

	"product/rpc/internal/svc"
	"product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductsByCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsByCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsByCategoryLogic {
	return &ProductsByCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsByCategoryLogic) ProductsByCategory(in *product.ProductsByCategoryRequest) (*product.ProductsByCategoryResponse, error) {
	// todo: add your logic here and delete this line

	return &product.ProductsByCategoryResponse{}, nil
}
