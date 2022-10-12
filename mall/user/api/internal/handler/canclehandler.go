package handler

import (
	"net/http"

	"go-zero-demo2/mall/user/api/internal/logic"
	"go-zero-demo2/mall/user/api/internal/svc"
	"go-zero-demo2/mall/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func cancleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCancleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCancleLogic(r.Context(), svcCtx)
		resp, err := l.Cancle(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
