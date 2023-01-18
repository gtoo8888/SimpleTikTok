package logic

import (
	"context"

	"go-zero-demo/BaseInterface/internal/svc"
	"go-zero-demo/BaseInterface/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowListLogic) RelationFollowList(req *types.RelationFollowListHandlerRequest) (resp *types.RelationFollowListHandlerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
