package logic

import (
	"context"

	"tiktokPlus/user/rpc/internal/svc"
	"tiktokPlus/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *RegisterLogic) Register(in *user.LoginReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line

	return &user.RegisterResp{}, nil
}
