package logic

import (
	"context"

	"tiktokPlus/user/api/internal/svc"
	"tiktokPlus/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload() (resp *types.UserUploadResp, err error) {
	// todo: add your logic here and delete this line

	return
}
