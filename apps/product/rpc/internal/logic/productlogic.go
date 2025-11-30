package logic

import (
	"context"

	"product/model"
	"product/rpc/internal/svc"
	"product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLogic) Product(in *product.ProductItemRequest) (*product.ProductItem, error) {
	// 先查询，确保存在
	resp, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(in.ProductId))
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, status.Errorf(codes.NotFound, "Product not found")
		}
		logx.Errorf("Product find one error:%v", err)
		return nil, status.Errorf(codes.Internal, "Database error: %v", err)
	}

	var (
		category *model.ProductCategory
		images   []*model.Image
	)

	// 并发查询（可能会导致对数据库的连接数增加）
	err = mr.Finish(
		func() error {
			category, err = l.svcCtx.ProductCategoryModel.FindOne(l.ctx, resp.CategoryId)
			if err == sqlx.ErrNotFound {
				return nil // 未找到时，不终止其他的
			}
			return err
		},

		func() error {
			images, err = l.svcCtx.ImageModel.FindAllByProductId(l.ctx, uint64(in.ProductId))
			if err == sqlx.ErrNotFound {
				return nil
			}
			return err
		},
	)

	if category == nil {
		return nil, status.Errorf(codes.NotFound, "未找到对应类型记录")
	}

	imageItems := make([]*product.Image, len(images))
	for _, img := range images {
		imageItems = append(imageItems, &product.Image{
			Id:  int64(img.Id),
			Url: img.Url,
		})
	}

	return &product.ProductItem{
		Id:          int64(resp.Id),
		Name:        resp.Name,
		Description: resp.Description,
		Price:       resp.Price,
		Stock:       resp.Stock,
		Category:    category.Name,
		Status:      resp.Status,
		Images:      imageItems,
	}, nil
}
