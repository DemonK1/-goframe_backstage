// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsList is the golang structure of table goods_list for DAO operations like Where/Data.
type GoodsList struct {
	g.Meta       `orm:"table:goods_list, do:true"`
	Id           interface{} // 主键
	Pid          interface{} // 类目 id(主表)
	PicUrl       interface{} // 主图
	Title        interface{} // 标题
	Price        interface{} // 价格(展示用)
	Account      interface{} // 账号
	AccountType  interface{} // 账号类型 如苹果安卓qq微信 山东一区
	RecyclePrice interface{} // 回收价
	SellPrice    interface{} // 出售价
	Submitter    interface{} // 提交人
	Status       interface{} // 状态 1 未上架  2 上架中 3 已下架 4 其他
	Examine      interface{} // 审核 1 未审核 2 审核中 3 审核未通过 4 其他
	CreateTime   *gtime.Time //
	UpdateTime   *gtime.Time //
	DeleteTime   *gtime.Time //
}
