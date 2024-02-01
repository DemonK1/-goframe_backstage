package service

import (
	"01-practice/internal/dao"
	"01-practice/internal/model"
	"01-practice/internal/model/entity"
	"context"
)

type GetSkuSceneService struct{}

func NewGetSkuSceneService() *GetSkuSceneService {
	return &GetSkuSceneService{}
}

// GetList 获取分类列表
func (s *GetSkuSceneService) GetList(ctx context.Context, param *model.SkuSceneReq) (res []*model.SkuSceneRes, err error) {
	ent := make([]*entity.SkuScene, 0)
	if err = dao.SkuScene.Ctx(ctx).Scan(&ent); err != nil {
		return nil, err
	}
	res = make([]*model.SkuSceneRes, 0)
	for _, v := range ent {
		resDetail := new(model.SkuSceneRes)
		resDetail.Id = v.Id
		resDetail.Name = v.Name
		res = append(res, resDetail)
	}
	return res, nil
}
