// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package home

import (
	"net/http"

	"api/internal/logic/home"
	"api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 首页Banner
func HomeBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := home.NewHomeBannerLogic(r.Context(), svcCtx)
		resp, err := l.HomeBanner()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
