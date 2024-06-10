package cart

import "errors"

var (
	ErrItemAlreadyExistsInCart = errors.New("商品已经存在")
	ErrCountInvalid            = errors.New("数量不能为负值")
	ErrInvalidItems            = errors.New("商品不存在")
)
