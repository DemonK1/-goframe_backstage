// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContraryDetailDao is the data access object for table contrary_detail.
type ContraryDetailDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns ContraryDetailColumns // columns contains all the column names of Table for convenient usage.
}

// ContraryDetailColumns defines and stores column names for table contrary_detail.
type ContraryDetailColumns struct {
	Id                string // 主键
	Pid               string // 主表的 id
	AccountSource     string // 账号来源 1 自己注册 2 购买 3 其他
	Ladder            string // 天梯赛 1 可排位 2 不可排位 3 其他
	GameLv            string // 游戏等级
	EpicShard         string // 史诗碎片
	ReviveCount       string // 复活币数量
	Voucher           string // 点券(期限/系统不算)
	TowerDefence      string // 塔防
	EpicSuperiorCount string // 史诗套装数量
	EpicSuperior      string // 史诗套装
	EpicSuperiorSkin  string // 史诗套皮肤(不是全套不要选)
	EmperorCore       string // 天地核心
	Wingman           string // 僚机
	Superior          string // 套装
	Artifact          string // 神器
	BlastingSuit      string // 爆破装
	Sniper            string // 狙击枪
	ViceWep           string // 副武器
	NearWep           string // 近身武器
	Role              string // 角色
	RoleWings         string // 翅膀
	RoleMask          string // 面具
	Backpack          string // 手环/背包
	RoleRing          string // 戒指
	FellowCreatures   string // 伴生灵
	MountPet          string // 坐骑宠物
	Treasury          string // 神武宝库
	HotProperty       string // 热门皮肤
	BrightSpot        string // 亮点
	CreateTime        string //
	UpdateTime        string //
	DeleteTime        string //
}

// contraryDetailColumns holds the columns for table contrary_detail.
var contraryDetailColumns = ContraryDetailColumns{
	Id:                "id",
	Pid:               "pid",
	AccountSource:     "account_source",
	Ladder:            "ladder",
	GameLv:            "game_lv",
	EpicShard:         "epic_shard",
	ReviveCount:       "revive_count",
	Voucher:           "voucher",
	TowerDefence:      "tower_defence",
	EpicSuperiorCount: "epic_superior_count",
	EpicSuperior:      "epic_superior",
	EpicSuperiorSkin:  "epic_superior_skin",
	EmperorCore:       "emperor_core",
	Wingman:           "wingman",
	Superior:          "superior",
	Artifact:          "artifact",
	BlastingSuit:      "blasting_suit",
	Sniper:            "sniper",
	ViceWep:           "vice_wep",
	NearWep:           "near_wep",
	Role:              "role",
	RoleWings:         "role_wings",
	RoleMask:          "role_mask",
	Backpack:          "backpack",
	RoleRing:          "role_ring",
	FellowCreatures:   "fellow_creatures",
	MountPet:          "mount_pet",
	Treasury:          "treasury",
	HotProperty:       "hot_property",
	BrightSpot:        "bright_spot",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
	DeleteTime:        "delete_time",
}

// NewContraryDetailDao creates and returns a new DAO object for table data access.
func NewContraryDetailDao() *ContraryDetailDao {
	return &ContraryDetailDao{
		group:   "default",
		table:   "contrary_detail",
		columns: contraryDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ContraryDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ContraryDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ContraryDetailDao) Columns() ContraryDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ContraryDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ContraryDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ContraryDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
