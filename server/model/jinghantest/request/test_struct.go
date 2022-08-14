package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/jinghantest"
)

type TestStructSearch struct {
	jinghantest.TestStruct
	request.PageInfo
}
