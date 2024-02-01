// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// KingGunplayDetail is the golang structure for table king_gunplay_detail.
type KingGunplayDetail struct {
	Id                  int         `json:"id"                  ` // 主键
	Pid                 int         `json:"pid"                 ` // 主表的 id
	AccountSource       string      `json:"accountSource"       ` // 账号来源 1 自己注册 2 购买 3 其他
	DoCount             string      `json:"doCount"             ` // 钻石数量
	LvVip               string      `json:"lvVip"               ` // vip等级
	WepHeroCount        string      `json:"wepHeroCount"        ` // 英雄级武器数量
	ThrowWepHeroCount   string      `json:"throwWepHeroCount"   ` // 英雄级投掷武器数量
	KingWepCount        string      `json:"kingWepCount"        ` // 王者武器数量
	CothSuperior        string      `json:"cothSuperior"        ` // 至尊套
	Superior            string      `json:"superior"            ` // 套装
	KingWep             string      `json:"kingWep"             ` // 王者武器
	HotWep              string      `json:"hotWep"              ` // 热门武器
	Grenade             string      `json:"grenade"             ` // 手雷
	RareRole            string      `json:"rareRole"            ` // 稀有角色
	RareProperty        string      `json:"rareProperty"        ` // 稀有道具/皮肤
	KingMainAndThreeWep string      `json:"kingMainAndThreeWep" ` // 王者主武器和三国武器整体皮肤
	ContactWay          string      `json:"contactWay"          ` // 联系方式
	BrightSpot          string      `json:"brightSpot"          ` // 亮点
	CreateTime          *gtime.Time `json:"createTime"          ` //
	UpdateTime          *gtime.Time `json:"updateTime"          ` //
	DeleteTime          *gtime.Time `json:"deleteTime"          ` //
}
