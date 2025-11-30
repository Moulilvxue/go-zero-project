// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package home

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeProductCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 商品评价
func NewHomeProductCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeProductCommentsLogic {
	return &HomeProductCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeProductCommentsLogic) HomeProductComments(req *types.CommentListReq) (resp *types.CommentListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
