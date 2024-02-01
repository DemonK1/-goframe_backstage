// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EliteDetailDao is the data access object for table elite_detail.
type EliteDetailDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns EliteDetailColumns // columns contains all the column names of Table for convenient usage.
}

// EliteDetailColumns defines and stores column names for table elite_detail.
type EliteDetailColumns struct {
	Id              string // 主键
	Pid             string // 主表的 id
	AccountSource   string // 账号来源 1 自己注册 2 购买 3 其他
	AntiAddiction   string // 健康系统有无限制
	KingMark        string // 王牌印记
	GameLv          string // 当前段位
	HotValue        string // 热力值
	SkinCount       string // 套装数量
	PistolSkinCount string // 枪皮数量
	PistolVfxCount  string // 特效枪数量
	PistolSkin      string // 优质枪皮
	VehicleCount    string // 载具数量
	VehicleSkin     string // 载具皮肤
	PinkCount       string // 粉装数量
	Skin            string // 优质套装
	SkinFill        string // 优质套装补充
	Parachute       string // 优质降落伞
	Backpack        string // 优质整套背包
	Aircraft        string // 飞行器
	BrightSpot      string // 亮点
	CreateTime      string //
	UpdateTime      string //
	DeleteTime      string //
}

// eliteDetailColumns holds the columns for table elite_detail.
var eliteDetailColumns = EliteDetailColumns{
	Id:              "id",
	Pid:             "pid",
	AccountSource:   "account_source",
	AntiAddiction:   "anti_addiction",
	KingMark:        "king_mark",
	GameLv:          "game_lv",
	HotValue:        "hot_value",
	SkinCount:       "skin_count",
	PistolSkinCount: "pistol_skin_count",
	PistolVfxCount:  "pistol_vfx_count",
	PistolSkin:      "pistol_skin",
	VehicleCount:    "vehicle_count",
	VehicleSkin:     "vehicle_skin",
	PinkCount:       "pink_count",
	Skin:            "skin",
	SkinFill:        "skin_fill",
	Parachute:       "parachute",
	Backpack:        "backpack",
	Aircraft:        "aircraft",
	BrightSpot:      "bright_spot",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
	DeleteTime:      "delete_time",
}

// NewEliteDetailDao creates and returns a new DAO object for table data access.
func NewEliteDetailDao() *EliteDetailDao {
	return &EliteDetailDao{
		group:   "default",
		table:   "elite_detail",
		columns: eliteDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EliteDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EliteDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EliteDetailDao) Columns() EliteDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EliteDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EliteDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EliteDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
