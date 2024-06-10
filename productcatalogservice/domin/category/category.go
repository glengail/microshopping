package category

import "errors"

var (
	//商品分类已存在
	ErrCategoryExistWithName = errors.New("商品分类已存在")
)
