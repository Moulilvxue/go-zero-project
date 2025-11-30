// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"net/http"

	"api/internal/logic/user"
	"api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户注销
func UserLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserLogoutLogic(r.Context(), svcCtx)
		resp, err := l.UserLogout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
