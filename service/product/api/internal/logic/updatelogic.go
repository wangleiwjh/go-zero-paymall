package logic

import (
	"context"
	"go-zero-mall-myself/service/product/pb/product"

	"go-zero-mall-myself/service/product/api/internal/svc"
	"go-zero-mall-myself/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &product.UpdateRequest{
		Id:     req.Id,
		Name:   req.Name,
		Desc:   req.Desc,
		Stock:  req.Stock,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateResponse{}, nil
}
