// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoodsListDetailDao is the data access object for table goods_list_detail.
type GoodsListDetailDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns GoodsListDetailColumns // columns contains all the column names of Table for convenient usage.
}

// GoodsListDetailColumns defines and stores column names for table goods_list_detail.
type GoodsListDetailColumns struct {
	Id                string // 主键
	Pid               string // 主表的 id
	AccountSource     string // 账号来源 1 自己注册 2 购买 3 其他
	AccountType       string // 区服 大区如 北部战区
	AccountTypeDetail string // 小区服 如: 山东一区
	RealName          string // 实名/防沉迷
	Isladder          string // 是否可以排位/禁言
	GameLv            string // 排位等级 如:大师
	VCountExAstrict   string // 总V数量去掉期限V
	GrenadeRole       string // 扣除AK猩红+手雷+人物+限期后英雄武器数量
	GrenadeCount      string // 手雷数量
	KingWepCount      string // 王者武器数量
	BlackGoldCount    string // 炫金武器数量
	VoucherCount      string // CF点券数量
	KingWep           string // 王者武器
	Superior          string // 步枪
	Sniper            string // 狙击枪
	Biochemistry      string // 生化
	Pistol            string // 手枪
	NearWep           string // 近身武器
	Role              string // 角色
	Grenade           string // 手雷
	Sound             string // 音效卡
	Skin              string // 皮肤
	PentPty           string // 永久道具
	Anniversary       string // 周年套装
	BrightSpot        string // 亮点
	CreateTime        string //
	UpdateTime        string //
	DeleteTime        string //
}

// goodsListDetailColumns holds the columns for table goods_list_detail.
var goodsListDetailColumns = GoodsListDetailColumns{
	Id:                "id",
	Pid:               "pid",
	AccountSource:     "account_source",
	AccountType:       "account_type",
	AccountTypeDetail: "account_type_detail",
	RealName:          "realName",
	Isladder:          "isladder",
	GameLv:            "game_lv",
	VCountExAstrict:   "v_count_ex_astrict",
	GrenadeRole:       "grenade_role",
	GrenadeCount:      "grenade_count",
	KingWepCount:      "king_wep_count",
	BlackGoldCount:    "black_gold_count",
	VoucherCount:      "voucher_count",
	KingWep:           "king_wep",
	Superior:          "superior",
	Sniper:            "sniper",
	Biochemistry:      "biochemistry",
	Pistol:            "pistol",
	NearWep:           "near_wep",
	Role:              "role",
	Grenade:           "grenade",
	Sound:             "sound",
	Skin:              "skin",
	PentPty:           "pent_pty",
	Anniversary:       "anniversary",
	BrightSpot:        "bright_spot",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
	DeleteTime:        "delete_time",
}

// NewGoodsListDetailDao creates and returns a new DAO object for table data access.
func NewGoodsListDetailDao() *GoodsListDetailDao {
	return &GoodsListDetailDao{
		group:   "default",
		table:   "goods_list_detail",
		columns: goodsListDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GoodsListDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoodsListDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GoodsListDetailDao) Columns() GoodsListDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoodsListDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoodsListDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoodsListDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
