// 自动生成模板TestStruct
package jinghantest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TestStruct 结构体
type TestStruct struct {
	global.GVA_MODEL
	TestD string `json:"testD" form:"testD" gorm:"column:test_d;comment:;"`
}

// TableName TestStruct 表名
func (TestStruct) TableName() string {
	return "test_struct"
}
