package controller

import (
	"01-practice/internal/model"
	"01-practice/internal/service"
	"context"
)

type GetSkuSceneController struct{}

// GetList 获取分类列表
func (c *GetSkuSceneController) GetList(ctx context.Context, req *model.SkuSceneReq) (res []*model.SkuSceneRes, err error) {
	res, err = service.NewGetSkuSceneService().GetList(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
