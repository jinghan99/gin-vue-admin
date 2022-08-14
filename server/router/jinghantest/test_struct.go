package jinghantest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestStructRouter struct {
}

// InitTestStructRouter 初始化 TestStruct 路由信息
func (s *TestStructRouter) InitTestStructRouter(Router *gin.RouterGroup) {
	test_structRouter := Router.Group("test_struct").Use(middleware.OperationRecord())
	test_structRouterWithoutRecord := Router.Group("test_struct")
	var test_structApi = v1.ApiGroupApp.JinghantestApiGroup.TestStructApi
	{
		test_structRouter.POST("createTestStruct", test_structApi.CreateTestStruct)             // 新建TestStruct
		test_structRouter.DELETE("deleteTestStruct", test_structApi.DeleteTestStruct)           // 删除TestStruct
		test_structRouter.DELETE("deleteTestStructByIds", test_structApi.DeleteTestStructByIds) // 批量删除TestStruct
		test_structRouter.PUT("updateTestStruct", test_structApi.UpdateTestStruct)              // 更新TestStruct
	}
	{
		test_structRouterWithoutRecord.GET("findTestStruct", test_structApi.FindTestStruct)       // 根据ID获取TestStruct
		test_structRouterWithoutRecord.GET("getTestStructList", test_structApi.GetTestStructList) // 获取TestStruct列表
	}
}
