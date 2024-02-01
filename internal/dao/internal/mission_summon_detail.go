// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MissionSummonDetailDao is the data access object for table mission_summon_detail.
type MissionSummonDetailDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns MissionSummonDetailColumns // columns contains all the column names of Table for convenient usage.
}

// MissionSummonDetailColumns defines and stores column names for table mission_summon_detail.
type MissionSummonDetailColumns struct {
	Id                 string // 主键
	Pid                string // 主表的 id
	AccountSource      string // 账号来源 1 自己注册 2 购买 3 其他
	MythicWepCount     string // 神话武器数量
	LegendWepCount     string // 传说武器数量
	LegendRoleCount    string // 传说角色数量
	LegendVehicleCount string // 传说载具数量
	LegendHiddenWep    string // 传说典藏武器
	LegendGlowWep      string // 传说辉光武器
	LegendWep          string // 传说武器
	MythicRole         string // 神话角色
	LegendRole         string // 传说角色
	Vehicle            string // 载具
	BrightSpot         string // 亮点
	CreateTime         string //
	UpdateTime         string //
	DeleteTime         string //
}

// missionSummonDetailColumns holds the columns for table mission_summon_detail.
var missionSummonDetailColumns = MissionSummonDetailColumns{
	Id:                 "id",
	Pid:                "pid",
	AccountSource:      "account_source",
	MythicWepCount:     "mythic_wep_count",
	LegendWepCount:     "legend_wep_count",
	LegendRoleCount:    "legend_role_count",
	LegendVehicleCount: "legend_vehicle_count",
	LegendHiddenWep:    "legend_hidden_wep",
	LegendGlowWep:      "legend_glow_wep",
	LegendWep:          "legend_wep",
	MythicRole:         "mythic_role",
	LegendRole:         "legend_role",
	Vehicle:            "vehicle",
	BrightSpot:         "bright_spot",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
	DeleteTime:         "delete_time",
}

// NewMissionSummonDetailDao creates and returns a new DAO object for table data access.
func NewMissionSummonDetailDao() *MissionSummonDetailDao {
	return &MissionSummonDetailDao{
		group:   "default",
		table:   "mission_summon_detail",
		columns: missionSummonDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MissionSummonDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MissionSummonDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MissionSummonDetailDao) Columns() MissionSummonDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MissionSummonDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MissionSummonDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MissionSummonDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
