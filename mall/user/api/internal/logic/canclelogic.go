package logic

import (
	"context"

	"go-zero-demo2/mall/user/api/internal/svc"
	"go-zero-demo2/mall/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancleLogic {
	return &CancleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancleLogic) Cancle(req *types.UserCancleRequest) (resp *types.UserCancleReply, err error) {
	// todo: add your logic here and delete this line

	return
}
