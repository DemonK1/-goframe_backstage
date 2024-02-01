// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MissionSummonDetail is the golang structure of table mission_summon_detail for DAO operations like Where/Data.
type MissionSummonDetail struct {
	g.Meta             `orm:"table:mission_summon_detail, do:true"`
	Id                 interface{} // 主键
	Pid                interface{} // 主表的 id
	AccountSource      interface{} // 账号来源 1 自己注册 2 购买 3 其他
	MythicWepCount     interface{} // 神话武器数量
	LegendWepCount     interface{} // 传说武器数量
	LegendRoleCount    interface{} // 传说角色数量
	LegendVehicleCount interface{} // 传说载具数量
	LegendHiddenWep    interface{} // 传说典藏武器
	LegendGlowWep      interface{} // 传说辉光武器
	LegendWep          interface{} // 传说武器
	MythicRole         interface{} // 神话角色
	LegendRole         interface{} // 传说角色
	Vehicle            interface{} // 载具
	BrightSpot         interface{} // 亮点
	CreateTime         *gtime.Time //
	UpdateTime         *gtime.Time //
	DeleteTime         *gtime.Time //
}
