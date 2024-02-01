// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MissionSummonDetail is the golang structure for table mission_summon_detail.
type MissionSummonDetail struct {
	Id                 int         `json:"id"                 ` // 主键
	Pid                int         `json:"pid"                ` // 主表的 id
	AccountSource      string      `json:"accountSource"      ` // 账号来源 1 自己注册 2 购买 3 其他
	MythicWepCount     string      `json:"mythicWepCount"     ` // 神话武器数量
	LegendWepCount     string      `json:"legendWepCount"     ` // 传说武器数量
	LegendRoleCount    string      `json:"legendRoleCount"    ` // 传说角色数量
	LegendVehicleCount string      `json:"legendVehicleCount" ` // 传说载具数量
	LegendHiddenWep    string      `json:"legendHiddenWep"    ` // 传说典藏武器
	LegendGlowWep      string      `json:"legendGlowWep"      ` // 传说辉光武器
	LegendWep          string      `json:"legendWep"          ` // 传说武器
	MythicRole         string      `json:"mythicRole"         ` // 神话角色
	LegendRole         string      `json:"legendRole"         ` // 传说角色
	Vehicle            string      `json:"vehicle"            ` // 载具
	BrightSpot         string      `json:"brightSpot"         ` // 亮点
	CreateTime         *gtime.Time `json:"createTime"         ` //
	UpdateTime         *gtime.Time `json:"updateTime"         ` //
	DeleteTime         *gtime.Time `json:"deleteTime"         ` //
}
