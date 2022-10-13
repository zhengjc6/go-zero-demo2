package handler

import (
	"net/http"

	"go-zero-demo2/mall/user/api/internal/common/response"
	"go-zero-demo2/mall/user/api/internal/logic"
	"go-zero-demo2/mall/user/api/internal/svc"
	"go-zero-demo2/mall/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
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
		response.Response(w, resp, err)
	}
}
