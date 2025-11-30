// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package payment

import (
	"net/http"

	"api/internal/logic/payment"
	"api/internal/svc"
	"api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 发起支付
func InitiatePaymentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InitiatePaymentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := payment.NewInitiatePaymentLogic(r.Context(), svcCtx)
		resp, err := l.InitiatePayment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
