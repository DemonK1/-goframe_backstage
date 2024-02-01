package service

import (
	"01-practice/internal/consts"
	"01-practice/internal/dao"
	"01-practice/internal/model"
	"01-practice/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"time"
)

type GoodsListService struct{}

func NewGoodsListService() *GoodsListService {
	return &GoodsListService{}
}

// GetList 查询商品列表
func (s *GoodsListService) GetList(ctx context.Context, param *model.GoodsListReq) (res []*model.GoodsListRes, err error) {
	m := dao.GoodsList.Ctx(ctx)
	var count int
	ent := make([]*entity.GoodsList, 0)
	if param.Page != 0 || param.Size != 0 {
		m = m.Page(param.Page, param.Size)
	}
	if param.Account != "" {
		m = m.Where(dao.GoodsList.Columns().Account, param.Account)
	}
	if param.RecyclePrice != "" {
		m = m.Where(dao.GoodsList.Columns().RecyclePrice, param.RecyclePrice)
	}
	if param.Submitter != "" {
		m = m.Where(dao.GoodsList.Columns().Submitter, param.Submitter)
	}
	if param.Title != "" {
		m = m.Where(dao.GoodsList.Columns().Title, param.Title)
	}
	if param.StartTime != "" || param.EndTime != "" {
		startTime, _ := time.Parse(consts.TimeFormat, param.StartTime)
		endTime, _ := time.Parse(consts.TimeFormat, param.EndTime)
		// 将开始时间往前推一秒，将结束时间往后推一天再减一秒
		startTime = startTime.Add(-1 * time.Second)
		endTime = endTime.Add(24 * time.Hour).Add(-1 * time.Second)
		m = m.WhereGTE(dao.GoodsList.Columns().CreateTime, startTime).WhereLTE(dao.GoodsList.Columns().CreateTime, endTime)
	}
	if err = m.WherePri(param.Id).ScanAndCount(&ent, &count, true); err != nil {
		return nil, gerror.New(consts.ServerErr)
	}
	if count == 0 {
		return nil, gerror.New(consts.QueryDataNull)
	}
	for _, v := range ent {
		resDetail := &model.GoodsListRes{
			CreateTime: v.CreateTime,
			Pid:        v.Pid,
			CreateGoodsReq: &model.CreateGoodsReq{
				Id:           v.Id,
				PicUrl:       v.PicUrl,
				Title:        v.Title,
				Price:        v.Price,
				Account:      v.Account,
				AccountType:  v.AccountType,
				RecyclePrice: v.RecyclePrice, // 回收价
				SellPrice:    v.SellPrice,    // 出售价
				Submitter:    v.Submitter,    // 提交人
			},
		}
		res = append(res, resDetail)
	}
	return res, nil
}

// GetGoodsById 查询商品详情
func (s *GoodsListService) GetGoodsById(ctx context.Context, param *model.IdReq) (res *model.GoodsDetailRes, err error) {
	res = new(model.GoodsDetailRes)
	if err = dao.GoodsListDetail.Ctx(ctx).WherePri(param.Id).Scan(&res); err != nil {
		return nil, gerror.New(consts.QueryDataNull)
	}
	return res, err
}

func (s *GoodsListService) Create(ctx context.Context, param *model.CreateGoodsReq) (res *model.CreateGoodsRes, err error) {
	goods := &entity.GoodsList{
		Pid:          param.Id,
		PicUrl:       param.PicUrl,
		Title:        param.Title,
		Price:        param.Price,
		Account:      param.Account,
		AccountType:  param.AccountType,
		RecyclePrice: param.RecyclePrice,
		SellPrice:    param.SellPrice,
		Submitter:    param.Submitter,
	}
	res = new(model.CreateGoodsRes)
	res.Id, err = dao.GoodsList.Ctx(ctx).Data(&goods).InsertAndGetId()
	if err != nil {
		return nil, gerror.New(consts.ServerErr)
	}
	return res, nil
}
