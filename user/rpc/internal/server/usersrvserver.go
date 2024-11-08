// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package server

import (
	"context"

	"tiktokPlus/user/rpc/internal/logic"
	"tiktokPlus/user/rpc/internal/svc"
	"tiktokPlus/user/rpc/pb/user"
)

type UserSrvServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserSrvServer
}

func NewUserSrvServer(svcCtx *svc.ServiceContext) *UserSrvServer {
	return &UserSrvServer{
		svcCtx: svcCtx,
	}
}

// 用户登录
func (s *UserSrvServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

// 用户注册
func (s *UserSrvServer) Register(ctx context.Context, in *user.LoginReq) (*user.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

// 获得用户信息
func (s *UserSrvServer) UserInfo(ctx context.Context, in *user.UserInfoReq) (*user.UserInfoResp, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

// 更改用户信息
func (s *UserSrvServer) UpdateInfo(ctx context.Context, in *user.UpdateReq) (*user.UpdateResp, error) {
	l := logic.NewUpdateInfoLogic(ctx, s.svcCtx)
	return l.UpdateInfo(in)
}
