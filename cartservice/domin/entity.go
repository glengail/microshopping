package cart

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint
	//可能用户与购物车之间是一对一或一对多的关系，并且需要在查询时直接获取用户信息，所以明确指定了外键引用
	//User user.User `gorm:"foreignKey:ID;references:UserID"` //指定字段User为外键，其中UserID依赖User表中的user_id(由gorm.model自动生成)
}

func NewCart(uid uint) *Cart {
	return &Cart{
		UserID: uid,
	}
}

type Item struct {
	gorm.Model
	/*
			可能是因为在查询时并不需要直接获取关联表（Product表和Cart表）的数据。
		或者可能是在实际业务逻辑中，并不需要通过外键引用来实现跨表查询，而是通过其他方式进行数据操作
	*/
	ProductSKU string
	CartID     uint
	Cart       Cart `gorm:"foreignKey:CartID" json:"-"`
	Count      int
}

func NewCartItem(productSKU string, cartID uint, count int) *Item {
	return &Item{
		ProductSKU: productSKU,
		CartID:     cartID,
		Count:      count,
	}

}
