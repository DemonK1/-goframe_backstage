// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EliteDetail is the golang structure of table elite_detail for DAO operations like Where/Data.
type EliteDetail struct {
	g.Meta          `orm:"table:elite_detail, do:true"`
	Id              interface{} // 主键
	Pid             interface{} // 主表的 id
	AccountSource   interface{} // 账号来源 1 自己注册 2 购买 3 其他
	AntiAddiction   interface{} // 健康系统有无限制
	KingMark        interface{} // 王牌印记
	GameLv          interface{} // 当前段位
	HotValue        interface{} // 热力值
	SkinCount       interface{} // 套装数量
	PistolSkinCount interface{} // 枪皮数量
	PistolVfxCount  interface{} // 特效枪数量
	PistolSkin      interface{} // 优质枪皮
	VehicleCount    interface{} // 载具数量
	VehicleSkin     interface{} // 载具皮肤
	PinkCount       interface{} // 粉装数量
	Skin            interface{} // 优质套装
	SkinFill        interface{} // 优质套装补充
	Parachute       interface{} // 优质降落伞
	Backpack        interface{} // 优质整套背包
	Aircraft        interface{} // 飞行器
	BrightSpot      interface{} // 亮点
	CreateTime      *gtime.Time //
	UpdateTime      *gtime.Time //
	DeleteTime      *gtime.Time //
}
