package pagination

import (
	"net/http"
	"strconv"

	pb "frontend/proto"

	"github.com/gin-gonic/gin"
)

var (
	//默认页数
	DefaultPageSize = 100
	//最大每页条数
	MaxPageSize = 1000
	//页数名称
	PageVar = "page"
	//每页条数名称
	PageSizeVar = "pageSize"
)

type Pages struct {
	Page       int         `json:"page"`       //当前页
	PageSize   int         `json:"pageSize"`   //页大小
	PageCount  int         `json:"pageCount"`  //页数
	TotalCount int         `json:"totalCount"` //总记录数
	Items      interface{} `json:"items"`      //记录切片
}

//实例化分页结构体

func NewPage(page int, pageSize int, totalCount int) *Pages {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	pageCount := -1
	//至少页数
	pageCount = (totalCount + pageSize - 1) / pageSize

	if totalCount > 0 {
		if page > pageCount {
			page = pageCount
		}
	}

	return &Pages{
		Page:       page,
		PageSize:   pageSize,
		PageCount:  pageCount,
		TotalCount: totalCount,
	}
}

// 根据http请求实例化分页结构体
func NewFromRequest1(req *http.Request, count int) *Pages {
	//获取当前页
	page := ParseInt(req.URL.Query().Get(PageVar), 1)
	//获取页大小
	pageSize := ParseInt(req.URL.Query().Get(PageSizeVar), DefaultPageSize)
	return NewPage(page, pageSize, count)

}

// 直接用pb的分页
func NewFromRequest(req *http.Request, count int) *pb.PageRequest {
	//获取当前页
	page := ParseInt(req.URL.Query().Get(PageVar), 1)
	//获取页大小
	pageSize := ParseInt(req.URL.Query().Get(PageSizeVar), DefaultPageSize)
	return &pb.PageRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
	}

}

// 根据gin请求实例化分页结构体
func NewFromGinRequest1(ctx *gin.Context, count int) *Pages {
	page := ParseInt(ctx.Query(PageVar), 1) //uri参数
	pageSize := ParseInt(ctx.Query(PageSizeVar), DefaultPageSize)
	return NewPage(page, pageSize, count)
}

// 直接用pb的分页
func NewFromGinRequest(ctx *gin.Context, count int) *pb.PageRequest {
	page := ParseInt(ctx.Query(PageVar), 1) //uri参数
	pageSize := ParseInt(ctx.Query(PageSizeVar), DefaultPageSize)
	return &pb.PageRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
	}
	//return NewPage(page, pageSize, count)
}

// offset
func (p *Pages) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// Limit
func (p *Pages) Limit() int {
	return p.PageSize
}
func ParseInt(str string, defaultValue int) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return value
}
