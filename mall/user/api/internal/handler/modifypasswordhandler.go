package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo2/mall/user/api/internal/logic"
	"go-zero-demo2/mall/user/api/internal/svc"
	"go-zero-demo2/mall/user/api/internal/types"
)

func modifypasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPasswordModifyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewModifypasswordLogic(r.Context(), svcCtx)
		resp, err := l.Modifypassword(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
