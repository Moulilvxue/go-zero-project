package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ImageModel = (*customImageModel)(nil)

var cacheImagePrefix = "cache:image:id:"

type (
	// ImageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customImageModel.
	ImageModel interface {
		imageModel
		FindAllByProductId(ctx context.Context, productId uint64) ([]*Image, error)
	}

	customImageModel struct {
		*defaultImageModel
	}
)

// NewImageModel returns a model for the database table.
func NewImageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ImageModel {
	return &customImageModel{
		defaultImageModel: newImageModel(conn, c, opts...),
	}
}

func (m *customImageModel) FindAllByProductId(ctx context.Context, productId uint64) ([]*Image, error) {
	key := fmt.Sprintf("%s%v", cacheImagePrefix, productId)
	var resp []*Image

	err := m.QueryRowCtx(ctx, &resp, key, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where product_id = ? ", imageRows, m.table)
		return conn.QueryRowsCtx(ctx, v, query, productId)
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
