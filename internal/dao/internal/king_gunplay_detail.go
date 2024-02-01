// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KingGunplayDetailDao is the data access object for table king_gunplay_detail.
type KingGunplayDetailDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns KingGunplayDetailColumns // columns contains all the column names of Table for convenient usage.
}

// KingGunplayDetailColumns defines and stores column names for table king_gunplay_detail.
type KingGunplayDetailColumns struct {
	Id                  string // 主键
	Pid                 string // 主表的 id
	AccountSource       string // 账号来源 1 自己注册 2 购买 3 其他
	DoCount             string // 钻石数量
	LvVip               string // vip等级
	WepHeroCount        string // 英雄级武器数量
	ThrowWepHeroCount   string // 英雄级投掷武器数量
	KingWepCount        string // 王者武器数量
	CothSuperior        string // 至尊套
	Superior            string // 套装
	KingWep             string // 王者武器
	HotWep              string // 热门武器
	Grenade             string // 手雷
	RareRole            string // 稀有角色
	RareProperty        string // 稀有道具/皮肤
	KingMainAndThreeWep string // 王者主武器和三国武器整体皮肤
	ContactWay          string // 联系方式
	BrightSpot          string // 亮点
	CreateTime          string //
	UpdateTime          string //
	DeleteTime          string //
}

// kingGunplayDetailColumns holds the columns for table king_gunplay_detail.
var kingGunplayDetailColumns = KingGunplayDetailColumns{
	Id:                  "id",
	Pid:                 "pid",
	AccountSource:       "account_source",
	DoCount:             "do_count",
	LvVip:               "lv_vip",
	WepHeroCount:        "wep_hero_count",
	ThrowWepHeroCount:   "throw_wep_hero_count",
	KingWepCount:        "king_wep_count",
	CothSuperior:        "coth_superior",
	Superior:            "superior",
	KingWep:             "king_wep",
	HotWep:              "hot_wep",
	Grenade:             "grenade",
	RareRole:            "rare_role",
	RareProperty:        "rare_property",
	KingMainAndThreeWep: "king_main_and_three_wep",
	ContactWay:          "contact_way",
	BrightSpot:          "bright_spot",
	CreateTime:          "create_time",
	UpdateTime:          "update_time",
	DeleteTime:          "delete_time",
}

// NewKingGunplayDetailDao creates and returns a new DAO object for table data access.
func NewKingGunplayDetailDao() *KingGunplayDetailDao {
	return &KingGunplayDetailDao{
		group:   "default",
		table:   "king_gunplay_detail",
		columns: kingGunplayDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *KingGunplayDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *KingGunplayDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *KingGunplayDetailDao) Columns() KingGunplayDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *KingGunplayDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *KingGunplayDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *KingGunplayDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
