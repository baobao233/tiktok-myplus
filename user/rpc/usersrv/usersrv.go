// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package usersrv

import (
	"context"

	"tiktokPlus/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginReq     = user.LoginReq
	LoginResp    = user.LoginResp
	RegisterReq  = user.RegisterReq
	RegisterResp = user.RegisterResp
	UpdateReq    = user.UpdateReq
	UpdateResp   = user.UpdateResp
	UserInfo     = user.UserInfo
	UserInfoReq  = user.UserInfoReq
	UserInfoResp = user.UserInfoResp

	UserSrv interface {
		// 用户登录
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		// 用户注册
		Register(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*RegisterResp, error)
		// 获得用户信息
		UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		// 更改用户信息
		UpdateInfo(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error)
	}

	defaultUserSrv struct {
		cli zrpc.Client
	}
)

func NewUserSrv(cli zrpc.Client) UserSrv {
	return &defaultUserSrv{
		cli: cli,
	}
}

// 用户登录
func (m *defaultUserSrv) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewUserSrvClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 用户注册
func (m *defaultUserSrv) Register(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := user.NewUserSrvClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

// 获得用户信息
func (m *defaultUserSrv) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := user.NewUserSrvClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

// 更改用户信息
func (m *defaultUserSrv) UpdateInfo(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error) {
	client := user.NewUserSrvClient(m.cli.Conn())
	return client.UpdateInfo(ctx, in, opts...)
}
