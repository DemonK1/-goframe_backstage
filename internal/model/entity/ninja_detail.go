// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NinjaDetail is the golang structure for table ninja_detail.
type NinjaDetail struct {
	Id            int         `json:"id"            ` // 主键
	Pid           int         `json:"pid"           ` // 主表的 id
	AccountSource string      `json:"accountSource" ` // 账号来源 1 自己注册 2 购买 3 其他
	AntiAddiction string      `json:"antiAddiction" ` // 健康系统有无限制
	GameLv        string      `json:"gameLv"        ` // 游戏等级
	VipLv         string      `json:"vipLv"         ` // vip等级
	Battle        string      `json:"battle"        ` // 战斗力
	SNinjaCount   string      `json:"sNinjaCount"   ` // 永久S忍者数量
	ANinjaCount   string      `json:"aNinjaCount"   ` // 永久A忍者数量
	SNinjaName    string      `json:"sNinjaName"    ` // 永久S忍者名称
	ANinjaName    string      `json:"aNinjaName"    ` // 永久A忍者名称
	BrightSpot    string      `json:"brightSpot"    ` // 亮点
	CreateTime    *gtime.Time `json:"createTime"    ` //
	UpdateTime    *gtime.Time `json:"updateTime"    ` //
	DeleteTime    *gtime.Time `json:"deleteTime"    ` //
}
