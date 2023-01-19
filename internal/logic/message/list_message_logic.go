package message

import (
	"context"

	"douyin/internal/svc"
	"douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMessageLogic {
	return &ListMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMessageLogic) ListMessage(req *types.ListMessageReq) (resp *types.SendMessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
