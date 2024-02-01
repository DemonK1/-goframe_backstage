// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// KingGunplayDetail is the golang structure of table king_gunplay_detail for DAO operations like Where/Data.
type KingGunplayDetail struct {
	g.Meta              `orm:"table:king_gunplay_detail, do:true"`
	Id                  interface{} // 主键
	Pid                 interface{} // 主表的 id
	AccountSource       interface{} // 账号来源 1 自己注册 2 购买 3 其他
	DoCount             interface{} // 钻石数量
	LvVip               interface{} // vip等级
	WepHeroCount        interface{} // 英雄级武器数量
	ThrowWepHeroCount   interface{} // 英雄级投掷武器数量
	KingWepCount        interface{} // 王者武器数量
	CothSuperior        interface{} // 至尊套
	Superior            interface{} // 套装
	KingWep             interface{} // 王者武器
	HotWep              interface{} // 热门武器
	Grenade             interface{} // 手雷
	RareRole            interface{} // 稀有角色
	RareProperty        interface{} // 稀有道具/皮肤
	KingMainAndThreeWep interface{} // 王者主武器和三国武器整体皮肤
	ContactWay          interface{} // 联系方式
	BrightSpot          interface{} // 亮点
	CreateTime          *gtime.Time //
	UpdateTime          *gtime.Time //
	DeleteTime          *gtime.Time //
}
