// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NinjaDetailDao is the data access object for table ninja_detail.
type NinjaDetailDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns NinjaDetailColumns // columns contains all the column names of Table for convenient usage.
}

// NinjaDetailColumns defines and stores column names for table ninja_detail.
type NinjaDetailColumns struct {
	Id            string // 主键
	Pid           string // 主表的 id
	AccountSource string // 账号来源 1 自己注册 2 购买 3 其他
	AntiAddiction string // 健康系统有无限制
	GameLv        string // 游戏等级
	VipLv         string // vip等级
	Battle        string // 战斗力
	SNinjaCount   string // 永久S忍者数量
	ANinjaCount   string // 永久A忍者数量
	SNinjaName    string // 永久S忍者名称
	ANinjaName    string // 永久A忍者名称
	BrightSpot    string // 亮点
	CreateTime    string //
	UpdateTime    string //
	DeleteTime    string //
}

// ninjaDetailColumns holds the columns for table ninja_detail.
var ninjaDetailColumns = NinjaDetailColumns{
	Id:            "id",
	Pid:           "pid",
	AccountSource: "account_source",
	AntiAddiction: "anti_addiction",
	GameLv:        "game_lv",
	VipLv:         "vip_lv",
	Battle:        "battle",
	SNinjaCount:   "s_ninja_count",
	ANinjaCount:   "a_ninja_count",
	SNinjaName:    "s_ninja_name",
	ANinjaName:    "a_ninja_name",
	BrightSpot:    "bright_spot",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	DeleteTime:    "delete_time",
}

// NewNinjaDetailDao creates and returns a new DAO object for table data access.
func NewNinjaDetailDao() *NinjaDetailDao {
	return &NinjaDetailDao{
		group:   "default",
		table:   "ninja_detail",
		columns: ninjaDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NinjaDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NinjaDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NinjaDetailDao) Columns() NinjaDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NinjaDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NinjaDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NinjaDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
