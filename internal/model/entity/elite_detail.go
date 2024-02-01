// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EliteDetail is the golang structure for table elite_detail.
type EliteDetail struct {
	Id              int         `json:"id"              ` // 主键
	Pid             int         `json:"pid"             ` // 主表的 id
	AccountSource   string      `json:"accountSource"   ` // 账号来源 1 自己注册 2 购买 3 其他
	AntiAddiction   string      `json:"antiAddiction"   ` // 健康系统有无限制
	KingMark        string      `json:"kingMark"        ` // 王牌印记
	GameLv          string      `json:"gameLv"          ` // 当前段位
	HotValue        string      `json:"hotValue"        ` // 热力值
	SkinCount       string      `json:"skinCount"       ` // 套装数量
	PistolSkinCount string      `json:"pistolSkinCount" ` // 枪皮数量
	PistolVfxCount  string      `json:"pistolVfxCount"  ` // 特效枪数量
	PistolSkin      string      `json:"pistolSkin"      ` // 优质枪皮
	VehicleCount    string      `json:"vehicleCount"    ` // 载具数量
	VehicleSkin     string      `json:"vehicleSkin"     ` // 载具皮肤
	PinkCount       string      `json:"pinkCount"       ` // 粉装数量
	Skin            string      `json:"skin"            ` // 优质套装
	SkinFill        string      `json:"skinFill"        ` // 优质套装补充
	Parachute       string      `json:"parachute"       ` // 优质降落伞
	Backpack        string      `json:"backpack"        ` // 优质整套背包
	Aircraft        string      `json:"aircraft"        ` // 飞行器
	BrightSpot      string      `json:"brightSpot"      ` // 亮点
	CreateTime      *gtime.Time `json:"createTime"      ` //
	UpdateTime      *gtime.Time `json:"updateTime"      ` //
	DeleteTime      *gtime.Time `json:"deleteTime"      ` //
}
