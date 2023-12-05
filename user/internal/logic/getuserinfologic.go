package logic

import (
	"context"
	"errors"
	"go_zero_study/user/internal/model"
	"go_zero_study/user/internal/svc"
	"go_zero_study/user/types/user"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息的方法
func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	u := &model.User{}
	logx.Debug(in.GetUserID())
	ret := svc.SqlProxy.First(u, in.UserID)
	if ret.Error != nil && !errors.Is(ret.Error, gorm.ErrRecordNotFound) {
		logx.Error(ret.Error)
		return nil, ret.Error
	}
	uInfo := &user.UserInfo{}
	if u != nil {
		uInfo.ID = u.ID
		uInfo.Name = u.Name
		uInfo.Hobbies = []string{"nv"}
	}
	return &user.GetUserInfoResponse{
		UserInfo: uInfo,
	}, nil
}
