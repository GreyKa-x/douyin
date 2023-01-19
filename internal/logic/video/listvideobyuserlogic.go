package video

import (
	"context"

	"douyin/internal/svc"
	"douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideoByUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListVideoByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideoByUserLogic {
	return &ListVideoByUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListVideoByUserLogic) ListVideoByUser(req *types.ListVideoByUserReq) (resp *types.ListVideoByUserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
