// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QqCardDetailDao is the data access object for table qq_card_detail.
type QqCardDetailDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns QqCardDetailColumns // columns contains all the column names of Table for convenient usage.
}

// QqCardDetailColumns defines and stores column names for table qq_card_detail.
type QqCardDetailColumns struct {
	Id              string // 主键
	Pid             string // 主表的 id
	AccountSource   string // 账号来源 1 自己注册 2 购买 3 其他
	AntiAddiction   string // 是否防沉迷 1 是 2 否
	LvVip           string // vip等级
	ACardCount      string // 永久A车数量
	TCardCount      string // 永久T车数量
	TCardSkin       string // 永久T车皮肤
	ACard           string // 永久A车
	SpecialPet      string // 特殊宠物
	HiddenMale      string // 典藏男
	HiddenFemale    string // 典藏女
	MagicMaleSkin   string // 男魔法套装
	MagicFemaleSkin string // 女魔法套装
	Mount           string // 坐骑
	BrightSpot      string // 亮点
	CreateTime      string //
	UpdateTime      string //
	DeleteTime      string //
}

// qqCardDetailColumns holds the columns for table qq_card_detail.
var qqCardDetailColumns = QqCardDetailColumns{
	Id:              "id",
	Pid:             "pid",
	AccountSource:   "account_source",
	AntiAddiction:   "anti_addiction",
	LvVip:           "lv_vip",
	ACardCount:      "a_card_count",
	TCardCount:      "t_card_count",
	TCardSkin:       "t_card_skin",
	ACard:           "a_card",
	SpecialPet:      "special_pet",
	HiddenMale:      "hidden_male",
	HiddenFemale:    "hidden_female",
	MagicMaleSkin:   "magic_male_skin",
	MagicFemaleSkin: "magic_female_skin",
	Mount:           "mount",
	BrightSpot:      "bright_spot",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
	DeleteTime:      "delete_time",
}

// NewQqCardDetailDao creates and returns a new DAO object for table data access.
func NewQqCardDetailDao() *QqCardDetailDao {
	return &QqCardDetailDao{
		group:   "default",
		table:   "qq_card_detail",
		columns: qqCardDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *QqCardDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *QqCardDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *QqCardDetailDao) Columns() QqCardDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *QqCardDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *QqCardDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *QqCardDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
