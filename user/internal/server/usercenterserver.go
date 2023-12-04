// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"go_zero_study/user/internal/logic"
	"go_zero_study/user/internal/svc"
	"go_zero_study/user/types/user"
)

type UserCenterServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserCenterServer
}

func NewUserCenterServer(svcCtx *svc.ServiceContext) *UserCenterServer {
	return &UserCenterServer{
		svcCtx: svcCtx,
	}
}

// 获取用户信息的方法
func (s *UserCenterServer) GetUserInfo(ctx context.Context, in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}