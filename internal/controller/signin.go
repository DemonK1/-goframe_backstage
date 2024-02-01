package controller

import (
	"01-practice/internal/consts"
	"01-practice/internal/model"
	"01-practice/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

type SignInController struct{}

func (c *SignInController) SignIn(ctx context.Context, req *model.UserSignInReq) (res *model.UserSignInRes, err error) {
	if len(req.Username) >= 15 || len(req.Username) <= 3 || len(req.Password) >= 16 || len(req.Password) <= 4 {
		return nil, gerror.New(consts.UsernameAndPasswordLen)
	}
	res, err = service.NewLoginLogic().SignIn(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
