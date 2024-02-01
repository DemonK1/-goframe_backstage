// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// QqCardDetail is the golang structure of table qq_card_detail for DAO operations like Where/Data.
type QqCardDetail struct {
	g.Meta          `orm:"table:qq_card_detail, do:true"`
	Id              interface{} // 主键
	Pid             interface{} // 主表的 id
	AccountSource   interface{} // 账号来源 1 自己注册 2 购买 3 其他
	AntiAddiction   interface{} // 是否防沉迷 1 是 2 否
	LvVip           interface{} // vip等级
	ACardCount      interface{} // 永久A车数量
	TCardCount      interface{} // 永久T车数量
	TCardSkin       interface{} // 永久T车皮肤
	ACard           interface{} // 永久A车
	SpecialPet      interface{} // 特殊宠物
	HiddenMale      interface{} // 典藏男
	HiddenFemale    interface{} // 典藏女
	MagicMaleSkin   interface{} // 男魔法套装
	MagicFemaleSkin interface{} // 女魔法套装
	Mount           interface{} // 坐骑
	BrightSpot      interface{} // 亮点
	CreateTime      *gtime.Time //
	UpdateTime      *gtime.Time //
	DeleteTime      *gtime.Time //
}
