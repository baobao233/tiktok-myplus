package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tiktokPlus/user/api/internal/logic"
	"tiktokPlus/user/api/internal/svc"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUpdateLogic(r.Context(), svcCtx)
		resp, err := l.Update()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
