// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KingHonourDetailDao is the data access object for table king_honour_detail.
type KingHonourDetailDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns KingHonourDetailColumns // columns contains all the column names of Table for convenient usage.
}

// KingHonourDetailColumns defines and stores column names for table king_honour_detail.
type KingHonourDetailColumns struct {
	Id               string // 主键
	Pid              string // 主表的 id
	AccountSource    string // 账号来源 1 自己注册 2 购买 3 其他
	GameLv           string // 当前段位
	HeroCount        string // 英雄数量
	SkinCount        string // 皮肤数量
	NoblemanLv       string // 贵族等级
	NoblemanPoints   string // 贵族积分
	InscriptionCount string // 五级铭文数量
	KingMark         string // 王者印记
	Sign             string // 国标
	PurpleMoney      string // 紫星币
	CrystalCount     string // 未兑换的水晶数量
	InsideSkin       string // 内测皮肤
	HiddenSkin       string // 荣耀典藏皮肤
	HotQualify       string // 热门限定
	BrightSpot       string // 亮点
	CreateTime       string //
	UpdateTime       string //
	DeleteTime       string //
}

// kingHonourDetailColumns holds the columns for table king_honour_detail.
var kingHonourDetailColumns = KingHonourDetailColumns{
	Id:               "id",
	Pid:              "pid",
	AccountSource:    "account_source",
	GameLv:           "game_lv",
	HeroCount:        "hero_count",
	SkinCount:        "skin_count",
	NoblemanLv:       "nobleman_lv",
	NoblemanPoints:   "nobleman_points",
	InscriptionCount: "inscription_count",
	KingMark:         "king_mark",
	Sign:             "sign",
	PurpleMoney:      "purple_money",
	CrystalCount:     "crystal_count",
	InsideSkin:       "inside_skin",
	HiddenSkin:       "hidden_skin",
	HotQualify:       "hot_qualify",
	BrightSpot:       "bright_spot",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
	DeleteTime:       "delete_time",
}

// NewKingHonourDetailDao creates and returns a new DAO object for table data access.
func NewKingHonourDetailDao() *KingHonourDetailDao {
	return &KingHonourDetailDao{
		group:   "default",
		table:   "king_honour_detail",
		columns: kingHonourDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *KingHonourDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *KingHonourDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *KingHonourDetailDao) Columns() KingHonourDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *KingHonourDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *KingHonourDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *KingHonourDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
