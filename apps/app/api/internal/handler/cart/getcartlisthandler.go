// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package cart

import (
	"net/http"

	"api/internal/logic/cart"
	"api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取购物车商品列表
func GetCartListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cart.NewGetCartListLogic(r.Context(), svcCtx)
		resp, err := l.GetCartList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
