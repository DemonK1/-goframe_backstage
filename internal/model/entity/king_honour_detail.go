// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// KingHonourDetail is the golang structure for table king_honour_detail.
type KingHonourDetail struct {
	Id               int         `json:"id"               ` // 主键
	Pid              int         `json:"pid"              ` // 主表的 id
	AccountSource    string      `json:"accountSource"    ` // 账号来源 1 自己注册 2 购买 3 其他
	GameLv           string      `json:"gameLv"           ` // 当前段位
	HeroCount        string      `json:"heroCount"        ` // 英雄数量
	SkinCount        string      `json:"skinCount"        ` // 皮肤数量
	NoblemanLv       string      `json:"noblemanLv"       ` // 贵族等级
	NoblemanPoints   string      `json:"noblemanPoints"   ` // 贵族积分
	InscriptionCount string      `json:"inscriptionCount" ` // 五级铭文数量
	KingMark         string      `json:"kingMark"         ` // 王者印记
	Sign             string      `json:"sign"             ` // 国标
	PurpleMoney      string      `json:"purpleMoney"      ` // 紫星币
	CrystalCount     string      `json:"crystalCount"     ` // 未兑换的水晶数量
	InsideSkin       string      `json:"insideSkin"       ` // 内测皮肤
	HiddenSkin       string      `json:"hiddenSkin"       ` // 荣耀典藏皮肤
	HotQualify       string      `json:"hotQualify"       ` // 热门限定
	BrightSpot       string      `json:"brightSpot"       ` // 亮点
	CreateTime       *gtime.Time `json:"createTime"       ` //
	UpdateTime       *gtime.Time `json:"updateTime"       ` //
	DeleteTime       *gtime.Time `json:"deleteTime"       ` //
}
