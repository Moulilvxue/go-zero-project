package svc

import (
	"product/model"
	"product/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config               config.Config
	ProductModel         model.ProductModel
	ProductCategoryModel model.ProductCategoryModel
	BannerModel          model.BannerModel
	ImageModel           model.ImageModel
	BizRedisClient       *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:               c,
		ProductModel:         model.NewProductModel(conn, c.CacheRedis),
		ProductCategoryModel: model.NewProductCategoryModel(conn, c.CacheRedis),
		BannerModel:          model.NewBannerModel(conn, c.CacheRedis),
		ImageModel:           model.NewImageModel(conn, c.CacheRedis),
		BizRedisClient:       redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
