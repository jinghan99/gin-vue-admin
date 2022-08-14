package jinghantest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/jinghantest"
	jinghantestReq "github.com/flipped-aurora/gin-vue-admin/server/model/jinghantest/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TestStructApi struct {
}

var test_structService = service.ServiceGroupApp.JinghantestServiceGroup.TestStructService

// CreateTestStruct 创建TestStruct
// @Tags TestStruct
// @Summary 创建TestStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body jinghantest.TestStruct true "创建TestStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test_struct/createTestStruct [post]
func (test_structApi *TestStructApi) CreateTestStruct(c *gin.Context) {
	var test_struct jinghantest.TestStruct
	_ = c.ShouldBindJSON(&test_struct)
	if err := test_structService.CreateTestStruct(test_struct); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestStruct 删除TestStruct
// @Tags TestStruct
// @Summary 删除TestStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body jinghantest.TestStruct true "删除TestStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /test_struct/deleteTestStruct [delete]
func (test_structApi *TestStructApi) DeleteTestStruct(c *gin.Context) {
	var test_struct jinghantest.TestStruct
	_ = c.ShouldBindJSON(&test_struct)
	if err := test_structService.DeleteTestStruct(test_struct); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestStructByIds 批量删除TestStruct
// @Tags TestStruct
// @Summary 批量删除TestStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /test_struct/deleteTestStructByIds [delete]
func (test_structApi *TestStructApi) DeleteTestStructByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := test_structService.DeleteTestStructByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestStruct 更新TestStruct
// @Tags TestStruct
// @Summary 更新TestStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body jinghantest.TestStruct true "更新TestStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /test_struct/updateTestStruct [put]
func (test_structApi *TestStructApi) UpdateTestStruct(c *gin.Context) {
	var test_struct jinghantest.TestStruct
	_ = c.ShouldBindJSON(&test_struct)
	if err := test_structService.UpdateTestStruct(test_struct); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestStruct 用id查询TestStruct
// @Tags TestStruct
// @Summary 用id查询TestStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query jinghantest.TestStruct true "用id查询TestStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /test_struct/findTestStruct [get]
func (test_structApi *TestStructApi) FindTestStruct(c *gin.Context) {
	var test_struct jinghantest.TestStruct
	_ = c.ShouldBindQuery(&test_struct)
	if retest_struct, err := test_structService.GetTestStruct(test_struct.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retest_struct": retest_struct}, c)
	}
}

// GetTestStructList 分页获取TestStruct列表
// @Tags TestStruct
// @Summary 分页获取TestStruct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query jinghantestReq.TestStructSearch true "分页获取TestStruct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test_struct/getTestStructList [get]
func (test_structApi *TestStructApi) GetTestStructList(c *gin.Context) {
	var pageInfo jinghantestReq.TestStructSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := test_structService.GetTestStructInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
