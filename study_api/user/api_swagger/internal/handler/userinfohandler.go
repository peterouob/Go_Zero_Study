package handler

import (
	"go_zero/common/response"
	"go_zero/study_api/user/api_jwt/internal/logic"
	"go_zero/study_api/user/api_jwt/internal/svc"
	"net/http"
)

func userinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo()
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
		response.Response(r, w, resp, err)
	}
}
