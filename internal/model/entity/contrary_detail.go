// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ContraryDetail is the golang structure for table contrary_detail.
type ContraryDetail struct {
	Id                int         `json:"id"                ` // 主键
	Pid               int         `json:"pid"               ` // 主表的 id
	AccountSource     string      `json:"accountSource"     ` // 账号来源 1 自己注册 2 购买 3 其他
	Ladder            string      `json:"ladder"            ` // 天梯赛 1 可排位 2 不可排位 3 其他
	GameLv            string      `json:"gameLv"            ` // 游戏等级
	EpicShard         string      `json:"epicShard"         ` // 史诗碎片
	ReviveCount       string      `json:"reviveCount"       ` // 复活币数量
	Voucher           string      `json:"voucher"           ` // 点券(期限/系统不算)
	TowerDefence      string      `json:"towerDefence"      ` // 塔防
	EpicSuperiorCount string      `json:"epicSuperiorCount" ` // 史诗套装数量
	EpicSuperior      string      `json:"epicSuperior"      ` // 史诗套装
	EpicSuperiorSkin  string      `json:"epicSuperiorSkin"  ` // 史诗套皮肤(不是全套不要选)
	EmperorCore       string      `json:"emperorCore"       ` // 天地核心
	Wingman           string      `json:"wingman"           ` // 僚机
	Superior          string      `json:"superior"          ` // 套装
	Artifact          string      `json:"artifact"          ` // 神器
	BlastingSuit      string      `json:"blastingSuit"      ` // 爆破装
	Sniper            string      `json:"sniper"            ` // 狙击枪
	ViceWep           string      `json:"viceWep"           ` // 副武器
	NearWep           string      `json:"nearWep"           ` // 近身武器
	Role              string      `json:"role"              ` // 角色
	RoleWings         string      `json:"roleWings"         ` // 翅膀
	RoleMask          string      `json:"roleMask"          ` // 面具
	Backpack          string      `json:"backpack"          ` // 手环/背包
	RoleRing          string      `json:"roleRing"          ` // 戒指
	FellowCreatures   string      `json:"fellowCreatures"   ` // 伴生灵
	MountPet          string      `json:"mountPet"          ` // 坐骑宠物
	Treasury          string      `json:"treasury"          ` // 神武宝库
	HotProperty       string      `json:"hotProperty"       ` // 热门皮肤
	BrightSpot        string      `json:"brightSpot"        ` // 亮点
	CreateTime        *gtime.Time `json:"createTime"        ` //
	UpdateTime        *gtime.Time `json:"updateTime"        ` //
	DeleteTime        *gtime.Time `json:"deleteTime"        ` //
}
