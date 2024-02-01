package controller

import (
	"01-practice/internal/model"
	"01-practice/internal/service"
	"context"
)

type GoodsListController struct{}

func (c *GoodsListController) GetList(ctx context.Context, req *model.GoodsListReq) (res []*model.GoodsListRes, err error) {
	res, err = service.NewGoodsListService().GetList(ctx, req)
	return res, err
}

func (c *GoodsListController) GetGoodsById(ctx context.Context, req *model.IdReq) (res *model.GoodsDetailRes, err error) {
	res, err = service.NewGoodsListService().GetGoodsById(ctx, req)
	return res, err
}

func (c *GoodsListController) Create(ctx context.Context, req *model.CreateGoodsReq) (res *model.CreateGoodsRes, err error) {
	res = new(model.CreateGoodsRes)
	res, err = service.NewGoodsListService().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, err
}
