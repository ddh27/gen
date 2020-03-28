package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type DeptManager struct {
	EmpNo    int         `gorm:"column:emp_no;primary_key" json:"emp_no"`
	DeptNo   null.String `gorm:"column:dept_no" json:"dept_no"`
	FromDate null.Time   `gorm:"column:from_date" json:"from_date"`
	ToDate   null.Time   `gorm:"column:to_date" json:"to_date"`
}

// TableName sets the insert table name for this struct type
func (d *DeptManager) TableName() string {
	return "dept_manager"
}
