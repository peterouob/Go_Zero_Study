package handler

import (
	"go_zero/common/response"
	"net/http"

	"go_zero/study_api/user/api_v2/internal/logic"
	"go_zero/study_api/user/api_v2/internal/svc"
)

func userinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo()
		response.Response(r, w, resp, err)
	}
}
