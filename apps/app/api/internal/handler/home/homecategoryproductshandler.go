// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package home

import (
	"net/http"

	"api/internal/logic/home"
	"api/internal/svc"
	"api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 分类商品列表
func HomeCategoryProductsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryProductsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := home.NewHomeCategoryProductsLogic(r.Context(), svcCtx)
		resp, err := l.HomeCategoryProducts(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
