// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// KingHonourDetail is the golang structure of table king_honour_detail for DAO operations like Where/Data.
type KingHonourDetail struct {
	g.Meta           `orm:"table:king_honour_detail, do:true"`
	Id               interface{} // 主键
	Pid              interface{} // 主表的 id
	AccountSource    interface{} // 账号来源 1 自己注册 2 购买 3 其他
	GameLv           interface{} // 当前段位
	HeroCount        interface{} // 英雄数量
	SkinCount        interface{} // 皮肤数量
	NoblemanLv       interface{} // 贵族等级
	NoblemanPoints   interface{} // 贵族积分
	InscriptionCount interface{} // 五级铭文数量
	KingMark         interface{} // 王者印记
	Sign             interface{} // 国标
	PurpleMoney      interface{} // 紫星币
	CrystalCount     interface{} // 未兑换的水晶数量
	InsideSkin       interface{} // 内测皮肤
	HiddenSkin       interface{} // 荣耀典藏皮肤
	HotQualify       interface{} // 热门限定
	BrightSpot       interface{} // 亮点
	CreateTime       *gtime.Time //
	UpdateTime       *gtime.Time //
	DeleteTime       *gtime.Time //
}
