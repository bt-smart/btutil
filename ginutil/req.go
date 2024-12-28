package ginutil

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetPageData 获取gin查询参数中的分页参数
func GetPageData(c *gin.Context) (int, int) {
	pn := c.DefaultQuery("pageNo", "1")
	ps := c.DefaultQuery("pageSize", "10")
	pageNo, err := strconv.Atoi(pn)
	if err != nil || pageNo < 1 {
		pageNo = 1
	}

	pageSize, err := strconv.Atoi(ps)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	return pageNo, pageSize
}

// GetIDFormParam 获取gin请求中路径参数中的uint64 的主键
func GetIDFormParam(c *gin.Context) uint64 {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		id = 0
	}
	return id
}

func GetQueryUint64ByString(c *gin.Context, query string) uint64 {
	str := c.Query(query)
	if str == "" {
		return 0
	}
	v, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		v = 0
	}
	return v
}
