// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoodsListDao is the data access object for table goods_list.
type GoodsListDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns GoodsListColumns // columns contains all the column names of Table for convenient usage.
}

// GoodsListColumns defines and stores column names for table goods_list.
type GoodsListColumns struct {
	Id           string // 主键
	Pid          string // 类目 id(主表)
	PicUrl       string // 主图
	Title        string // 标题
	Price        string // 价格(展示用)
	Account      string // 账号
	AccountType  string // 账号类型 如苹果安卓qq微信 山东一区
	RecyclePrice string // 回收价
	SellPrice    string // 出售价
	Submitter    string // 提交人
	Status       string // 状态 1 未上架  2 上架中 3 已下架 4 其他
	Examine      string // 审核 1 未审核 2 审核中 3 审核未通过 4 其他
	CreateTime   string //
	UpdateTime   string //
	DeleteTime   string //
}

// goodsListColumns holds the columns for table goods_list.
var goodsListColumns = GoodsListColumns{
	Id:           "id",
	Pid:          "pid",
	PicUrl:       "pic_url",
	Title:        "title",
	Price:        "price",
	Account:      "account",
	AccountType:  "account_type",
	RecyclePrice: "recycle_price",
	SellPrice:    "sell_price",
	Submitter:    "submitter",
	Status:       "status",
	Examine:      "examine",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
	DeleteTime:   "delete_time",
}

// NewGoodsListDao creates and returns a new DAO object for table data access.
func NewGoodsListDao() *GoodsListDao {
	return &GoodsListDao{
		group:   "default",
		table:   "goods_list",
		columns: goodsListColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GoodsListDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoodsListDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GoodsListDao) Columns() GoodsListColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoodsListDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoodsListDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoodsListDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
