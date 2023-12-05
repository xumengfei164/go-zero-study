package logic

import (
	"context"
	"go_zero_study/search/internal/svc"
	"go_zero_study/search/internal/types"
	"go_zero_study/user/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchInfoLogic {
	return &SearchInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchInfoLogic) SearchInfo(req *types.SearchReq) (resp *types.SearchReply, err error) {
	//l.svcCtx.UserCenterServer = usercenter.NewUserCenter()
	ret, err := l.svcCtx.UC.GetUserInfo(
		l.ctx,
		&usercenter.GetUserInfoRequest{
			UserID: req.Id,
		})
	if err != nil {
		logx.Error("search:", err)
		return nil, err
	}
	rs := &types.SearchReply{}
	if ret.UserInfo != nil {
		rs.Hobbies = ret.UserInfo.GetHobbies()
		rs.Name = ret.UserInfo.GetName()
		rs.Id = ret.UserInfo.GetID()

	}
	return rs, nil
}
