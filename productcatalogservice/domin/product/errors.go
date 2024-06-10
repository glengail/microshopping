package product

import "errors"

var (
	ErrProductNotFound         = errors.New("商品不存在")
	ErrProductStockIsNotEnough = errors.New("商品库存不足")
)
