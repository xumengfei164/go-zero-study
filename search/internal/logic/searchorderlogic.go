package logic

import (
	"context"
	"go_zero_study/search/internal/model"
	"go_zero_study/search/internal/svc"
	"go_zero_study/search/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOrderLogic {
	return &SearchOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchOrderLogic) SearchOrder(req *types.SearchOrderReq) (resp *types.SearchOrderReply, err error) {
	//data := &model.Orders{
	//	Id:   1,
	//	Name: "phone",
	//	Total: sql.NullFloat64{
	//		Float64: 13444.5,
	//		Valid:   true,
	//	},
	//	CreateTime: 1701767029,
	//	UserID:     1,
	//}
	//_, err = l.svcCtx.Model.Insert(l.ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	ret, err := l.svcCtx.Model.FindOne(l.ctx, int64(req.Id))
	if err != nil && err != model.ErrNotFound {
		logx.Error("find:", err)
		return nil, err
	}
	reply := &types.SearchOrderReply{}
	if ret != nil {
		reply.Name = ret.Name
		reply.Hobbies = []string{}
	}
	return reply, nil
}
