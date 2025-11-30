package logic

import (
	"context"

	"product/rpc/internal/svc"
	"product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GetBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBannerLogic {
	return &GetBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBannerLogic) GetBanner(in *product.BannerRequest) (*product.BannerResponse, error) {
	resp, err := l.svcCtx.BannerModel.FindAllBanners(l.ctx)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return &product.BannerResponse{}, nil
		}

		logx.Errorf("Banner find all error:%v", err)
		return nil, err
	}

	// 拼装返回结果

	return &product.BannerResponse{}, nil
}
