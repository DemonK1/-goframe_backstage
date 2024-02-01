package test

import (
	"01-practice/internal/model"
	"01-practice/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestCreate(t *testing.T) {
	svc := service.NewGoodsListService()
	param := &model.CreateGoodsReq{
		Id:          0,
		PicUrl:      "https://image.url",
		Name:        "创建王者荣耀商品的标题",
		Price:       "创建王者荣耀商品的价格",
		AccountType: "苹果qq",
	}
	create, err := svc.Create(gctx.New(), param)
	if err != nil {
		fmt.Println(err)
	}
	jm, _ := json.Marshal(create)
	fmt.Println(jm)
}
