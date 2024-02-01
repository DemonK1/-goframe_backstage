// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsListDetail is the golang structure for table goods_list_detail.
type GoodsListDetail struct {
	Id                int         `json:"id"                ` // 主键
	Pid               int         `json:"pid"               ` // 主表的 id
	AccountSource     string      `json:"accountSource"     ` // 账号来源 1 自己注册 2 购买 3 其他
	AccountType       string      `json:"accountType"       ` // 区服 大区如 北部战区
	AccountTypeDetail string      `json:"accountTypeDetail" ` // 小区服 如: 山东一区
	RealName          string      `json:"realName"          ` // 实名/防沉迷
	Isladder          string      `json:"isladder"          ` // 是否可以排位/禁言
	GameLv            string      `json:"gameLv"            ` // 排位等级 如:大师
	VCountExAstrict   string      `json:"vCountExAstrict"   ` // 总V数量去掉期限V
	GrenadeRole       string      `json:"grenadeRole"       ` // 扣除AK猩红+手雷+人物+限期后英雄武器数量
	GrenadeCount      string      `json:"grenadeCount"      ` // 手雷数量
	KingWepCount      string      `json:"kingWepCount"      ` // 王者武器数量
	BlackGoldCount    string      `json:"blackGoldCount"    ` // 炫金武器数量
	VoucherCount      string      `json:"voucherCount"      ` // CF点券数量
	KingWep           string      `json:"kingWep"           ` // 王者武器
	Superior          string      `json:"superior"          ` // 步枪
	Sniper            string      `json:"sniper"            ` // 狙击枪
	Biochemistry      string      `json:"biochemistry"      ` // 生化
	Pistol            string      `json:"pistol"            ` // 手枪
	NearWep           string      `json:"nearWep"           ` // 近身武器
	Role              string      `json:"role"              ` // 角色
	Grenade           string      `json:"grenade"           ` // 手雷
	Sound             string      `json:"sound"             ` // 音效卡
	Skin              string      `json:"skin"              ` // 皮肤
	PentPty           string      `json:"pentPty"           ` // 永久道具
	Anniversary       string      `json:"anniversary"       ` // 周年套装
	BrightSpot        string      `json:"brightSpot"        ` // 亮点
	CreateTime        *gtime.Time `json:"createTime"        ` //
	UpdateTime        *gtime.Time `json:"updateTime"        ` //
	DeleteTime        *gtime.Time `json:"deleteTime"        ` //
}
