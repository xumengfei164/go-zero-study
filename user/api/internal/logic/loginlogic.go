package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"go_zero_study/user/api/internal/svc"
	"go_zero_study/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	userID := req.Id
	userPass := req.Password
	info, err := l.svcCtx.Model.FindOne(l.ctx, int64(userID))
	if err != nil {
		logx.Error(err)
		return nil, errors.New("未找到有效用户")
	}
	if info.Password != userPass {
		logx.Error(info)
		return nil, errors.New("密码错误")
	}
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, int64(userID))
	if err != nil {
		logx.Error(err)
		return nil, errors.New("生成token失败")
	}
	rsp := &types.LoginReply{
		SecretKey: token,
	}
	return rsp, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
