package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero/common/response"
	"go_zero/study_api/user/api_jwt/internal/logic"
	"go_zero/study_api/user/api_jwt/internal/svc"
	"go_zero/study_api/user/api_jwt/internal/types"
	"net/http"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
		response.Response(r, w, resp, err)
	}
}
