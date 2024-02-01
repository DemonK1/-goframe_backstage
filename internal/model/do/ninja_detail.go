// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NinjaDetail is the golang structure of table ninja_detail for DAO operations like Where/Data.
type NinjaDetail struct {
	g.Meta        `orm:"table:ninja_detail, do:true"`
	Id            interface{} // 主键
	Pid           interface{} // 主表的 id
	AccountSource interface{} // 账号来源 1 自己注册 2 购买 3 其他
	AntiAddiction interface{} // 健康系统有无限制
	GameLv        interface{} // 游戏等级
	VipLv         interface{} // vip等级
	Battle        interface{} // 战斗力
	SNinjaCount   interface{} // 永久S忍者数量
	ANinjaCount   interface{} // 永久A忍者数量
	SNinjaName    interface{} // 永久S忍者名称
	ANinjaName    interface{} // 永久A忍者名称
	BrightSpot    interface{} // 亮点
	CreateTime    *gtime.Time //
	UpdateTime    *gtime.Time //
	DeleteTime    *gtime.Time //
}
