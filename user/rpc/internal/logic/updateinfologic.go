package logic

import (
	"context"

	"tiktokPlus/user/rpc/internal/svc"
	"tiktokPlus/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更改用户信息
func (l *UpdateInfoLogic) UpdateInfo(in *user.UpdateReq) (*user.UpdateResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateResp{}, nil
}
