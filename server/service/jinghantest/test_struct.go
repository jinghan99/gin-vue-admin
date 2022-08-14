package jinghantest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/jinghantest"
	jinghantestReq "github.com/flipped-aurora/gin-vue-admin/server/model/jinghantest/request"
)

type TestStructService struct {
}

// CreateTestStruct 创建TestStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_structService *TestStructService) CreateTestStruct(test_struct jinghantest.TestStruct) (err error) {
	err = global.GVA_DB.Create(&test_struct).Error
	return err
}

// DeleteTestStruct 删除TestStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_structService *TestStructService) DeleteTestStruct(test_struct jinghantest.TestStruct) (err error) {
	err = global.GVA_DB.Delete(&test_struct).Error
	return err
}

// DeleteTestStructByIds 批量删除TestStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_structService *TestStructService) DeleteTestStructByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]jinghantest.TestStruct{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTestStruct 更新TestStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_structService *TestStructService) UpdateTestStruct(test_struct jinghantest.TestStruct) (err error) {
	err = global.GVA_DB.Save(&test_struct).Error
	return err
}

// GetTestStruct 根据id获取TestStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_structService *TestStructService) GetTestStruct(id uint) (test_struct jinghantest.TestStruct, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&test_struct).Error
	return
}

// GetTestStructInfoList 分页获取TestStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_structService *TestStructService) GetTestStructInfoList(info jinghantestReq.TestStructSearch) (list []jinghantest.TestStruct, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&jinghantest.TestStruct{})
	var test_structs []jinghantest.TestStruct
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&test_structs).Error
	return test_structs, total, err
}
