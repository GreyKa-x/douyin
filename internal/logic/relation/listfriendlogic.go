package relation

import (
	"context"

	"douyin/internal/svc"
	"douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFriendLogic {
	return &ListFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFriendLogic) ListFriend(req *types.ListFriendReq) (resp *types.ListFriendResp, err error) {
	// todo: add your logic here and delete this line

	return
}
