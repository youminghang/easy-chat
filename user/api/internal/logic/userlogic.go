package logic

import (
	"context"

	"github.com/youminghang/easy-chat/user/api/internal/svc"
	"github.com/youminghang/easy-chat/user/api/internal/types"
	"github.com/youminghang/easy-chat/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line
	getUserResp, err := l.svcCtx.GetUser(l.ctx, &userclient.GetUserReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserResp{
		Id:    getUserResp.Id,
		Name:  getUserResp.Name,
		Phone: getUserResp.Phone,
	}, nil
}
