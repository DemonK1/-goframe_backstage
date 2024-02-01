// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// QqCardDetail is the golang structure for table qq_card_detail.
type QqCardDetail struct {
	Id              int         `json:"id"              ` // 主键
	Pid             int         `json:"pid"             ` // 主表的 id
	AccountSource   string      `json:"accountSource"   ` // 账号来源 1 自己注册 2 购买 3 其他
	AntiAddiction   string      `json:"antiAddiction"   ` // 是否防沉迷 1 是 2 否
	LvVip           string      `json:"lvVip"           ` // vip等级
	ACardCount      string      `json:"aCardCount"      ` // 永久A车数量
	TCardCount      string      `json:"tCardCount"      ` // 永久T车数量
	TCardSkin       string      `json:"tCardSkin"       ` // 永久T车皮肤
	ACard           string      `json:"aCard"           ` // 永久A车
	SpecialPet      string      `json:"specialPet"      ` // 特殊宠物
	HiddenMale      string      `json:"hiddenMale"      ` // 典藏男
	HiddenFemale    string      `json:"hiddenFemale"    ` // 典藏女
	MagicMaleSkin   string      `json:"magicMaleSkin"   ` // 男魔法套装
	MagicFemaleSkin string      `json:"magicFemaleSkin" ` // 女魔法套装
	Mount           string      `json:"mount"           ` // 坐骑
	BrightSpot      string      `json:"brightSpot"      ` // 亮点
	CreateTime      *gtime.Time `json:"createTime"      ` //
	UpdateTime      *gtime.Time `json:"updateTime"      ` //
	DeleteTime      *gtime.Time `json:"deleteTime"      ` //
}
