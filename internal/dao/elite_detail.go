// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"01-practice/internal/dao/internal"
)

// internalEliteDetailDao is internal type for wrapping internal DAO implements.
type internalEliteDetailDao = *internal.EliteDetailDao

// eliteDetailDao is the data access object for table elite_detail.
// You can define custom methods on it to extend its functionality as you wish.
type eliteDetailDao struct {
	internalEliteDetailDao
}

var (
	// EliteDetail is globally public accessible object for table elite_detail operations.
	EliteDetail = eliteDetailDao{
		internal.NewEliteDetailDao(),
	}
)

// Fill with you ideas below.
