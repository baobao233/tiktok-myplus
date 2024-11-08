package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tiktokPlus/user/api/internal/logic"
	"tiktokPlus/user/api/internal/svc"
	"tiktokPlus/user/api/internal/types"
)

// 用户管理
func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
