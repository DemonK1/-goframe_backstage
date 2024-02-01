// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsListDetail is the golang structure of table goods_list_detail for DAO operations like Where/Data.
type GoodsListDetail struct {
	g.Meta            `orm:"table:goods_list_detail, do:true"`
	Id                interface{} // 主键
	Pid               interface{} // 主表的 id
	AccountSource     interface{} // 账号来源 1 自己注册 2 购买 3 其他
	AccountType       interface{} // 区服 大区如 北部战区
	AccountTypeDetail interface{} // 小区服 如: 山东一区
	RealName          interface{} // 实名/防沉迷
	Isladder          interface{} // 是否可以排位/禁言
	GameLv            interface{} // 排位等级 如:大师
	VCountExAstrict   interface{} // 总V数量去掉期限V
	GrenadeRole       interface{} // 扣除AK猩红+手雷+人物+限期后英雄武器数量
	GrenadeCount      interface{} // 手雷数量
	KingWepCount      interface{} // 王者武器数量
	BlackGoldCount    interface{} // 炫金武器数量
	VoucherCount      interface{} // CF点券数量
	KingWep           interface{} // 王者武器
	Superior          interface{} // 步枪
	Sniper            interface{} // 狙击枪
	Biochemistry      interface{} // 生化
	Pistol            interface{} // 手枪
	NearWep           interface{} // 近身武器
	Role              interface{} // 角色
	Grenade           interface{} // 手雷
	Sound             interface{} // 音效卡
	Skin              interface{} // 皮肤
	PentPty           interface{} // 永久道具
	Anniversary       interface{} // 周年套装
	BrightSpot        interface{} // 亮点
	CreateTime        *gtime.Time //
	UpdateTime        *gtime.Time //
	DeleteTime        *gtime.Time //
}
