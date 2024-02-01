// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ContraryDetail is the golang structure of table contrary_detail for DAO operations like Where/Data.
type ContraryDetail struct {
	g.Meta            `orm:"table:contrary_detail, do:true"`
	Id                interface{} // 主键
	Pid               interface{} // 主表的 id
	AccountSource     interface{} // 账号来源 1 自己注册 2 购买 3 其他
	Ladder            interface{} // 天梯赛 1 可排位 2 不可排位 3 其他
	GameLv            interface{} // 游戏等级
	EpicShard         interface{} // 史诗碎片
	ReviveCount       interface{} // 复活币数量
	Voucher           interface{} // 点券(期限/系统不算)
	TowerDefence      interface{} // 塔防
	EpicSuperiorCount interface{} // 史诗套装数量
	EpicSuperior      interface{} // 史诗套装
	EpicSuperiorSkin  interface{} // 史诗套皮肤(不是全套不要选)
	EmperorCore       interface{} // 天地核心
	Wingman           interface{} // 僚机
	Superior          interface{} // 套装
	Artifact          interface{} // 神器
	BlastingSuit      interface{} // 爆破装
	Sniper            interface{} // 狙击枪
	ViceWep           interface{} // 副武器
	NearWep           interface{} // 近身武器
	Role              interface{} // 角色
	RoleWings         interface{} // 翅膀
	RoleMask          interface{} // 面具
	Backpack          interface{} // 手环/背包
	RoleRing          interface{} // 戒指
	FellowCreatures   interface{} // 伴生灵
	MountPet          interface{} // 坐骑宠物
	Treasury          interface{} // 神武宝库
	HotProperty       interface{} // 热门皮肤
	BrightSpot        interface{} // 亮点
	CreateTime        *gtime.Time //
	UpdateTime        *gtime.Time //
	DeleteTime        *gtime.Time //
}
