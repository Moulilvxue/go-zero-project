package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BannerModel = (*customBannerModel)(nil)
var status = 1
var cacheBannerPrefix = "cache:banner:list"

type (
	// BannerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBannerModel.
	BannerModel interface {
		bannerModel
		FindAllBanners(ctx context.Context) ([]*Banner, error)
	}

	customBannerModel struct {
		*defaultBannerModel
	}
)

// NewBannerModel returns a model for the database table.
func NewBannerModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) BannerModel {
	return &customBannerModel{
		defaultBannerModel: newBannerModel(conn, c, opts...),
	}
}

func (m *customBannerModel) FindAllBanners(ctx context.Context) ([]*Banner, error) {
	var resp []*Banner
	err := m.QueryRowCtx(ctx, &resp, cacheBannerPrefix, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where status = ?", bannerRows, m.table)
		return conn.QueryRowsCtx(ctx, v, query, status)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, ErrNotFound
	default:
		return nil, err
	}
}
