package logic

import (
	"context"
	"order/internal/svc"
	"order/internal/types"
	"user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line
	userId := l.getOrderById(req.Id)
	// 根据用户id去user服务获取用户信息
	userResponse, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}
	return &types.OrderReply{
		Id:       req.Id,
		Name:     "hello order name",
		UserName: userResponse.Name,
	}, nil
}

func (l *GetOrderLogic) getOrderById(id string) string {
	return "1"
}
