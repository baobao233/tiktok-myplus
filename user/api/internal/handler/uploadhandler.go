package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tiktokPlus/user/api/internal/logic"
	"tiktokPlus/user/api/internal/svc"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
