package relation

import (
	"context"

	"douyin/internal/svc"
	"douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowLogic {
	return &ListFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFollowLogic) ListFollow(req *types.ListFollowReq) (resp *types.ListFollowResp, err error) {
	// todo: add your logic here and delete this line

	return
}
