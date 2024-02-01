package service

import (
	"01-practice/internal/commonFunc"
	"01-practice/internal/consts"
	"01-practice/internal/dao"
	"01-practice/internal/model"
	"01-practice/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/errors/gerror"
)

type LoginLogic struct{}

func NewLoginLogic() *LoginLogic {
	svc := new(LoginLogic)
	return svc
}

func (s *LoginLogic) SignIn(ctx context.Context, req *model.UserSignInReq) (res *model.UserSignInRes, err error) {
	count, err := dao.Users.Ctx(ctx).Fields(dao.Users.Columns().Username).Where(dao.Users.Columns().Username, req.Username).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New(consts.UserExist)
	}
	data := new(entity.Users)
	newPsw := commonFunc.EncryptPassword(req.Password)
	data.Username = req.Username
	data.Password = newPsw
	data.Nickname = req.Nickname
	// 开启事务 error 为 nil 执行 commit 操作 否则执行 RollBack 操作
	err = g.DB().Transaction(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
		res = new(model.UserSignInRes)
		id, err := dao.Users.Ctx(ctx).FieldsEx(consts.CreateFieldEx...).Data(data).InsertAndGetId()
		if err != nil {
			return err
		}
		if err = dao.Users.Ctx(ctx).WherePri(id).FieldsEx(dao.Users.Columns().Password).Scan(&res); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, gerror.New(consts.ServerErr)
	}
	return res, nil
}
