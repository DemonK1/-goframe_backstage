// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsList is the golang structure for table goods_list.
type GoodsList struct {
	Id           int         `json:"id"           ` // 主键
	Pid          int         `json:"pid"          ` // 类目 id(主表)
	PicUrl       string      `json:"picUrl"       ` // 主图
	Title        string      `json:"title"        ` // 标题
	Price        string      `json:"price"        ` // 价格(展示用)
	Account      string      `json:"account"      ` // 账号
	AccountType  string      `json:"accountType"  ` // 账号类型 如苹果安卓qq微信 山东一区
	RecyclePrice string      `json:"recyclePrice" ` // 回收价
	SellPrice    string      `json:"sellPrice"    ` // 出售价
	Submitter    string      `json:"submitter"    ` // 提交人
	Status       string      `json:"status"       ` // 状态 1 未上架  2 上架中 3 已下架 4 其他
	Examine      string      `json:"examine"      ` // 审核 1 未审核 2 审核中 3 审核未通过 4 其他
	CreateTime   *gtime.Time `json:"createTime"   ` //
	UpdateTime   *gtime.Time `json:"updateTime"   ` //
	DeleteTime   *gtime.Time `json:"deleteTime"   ` //
}
