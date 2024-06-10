package cart

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
func (r *Repository) Migration() {
	err := r.DB.AutoMigrate(&Cart{})
	if err != nil {
		log.Println(err)
	}
}

// 更新
func (r *Repository) Update(cart *Cart) error {
	err := r.DB.Save(cart).Error
	if err != nil {
		return err
	}
	return nil
}

// 硬删除
func (r *Repository) Delete(cart *Cart) error {
	err := r.DB.Where(&cart).Delete(&Cart{}).Error
	if err != nil {
		return err
	}
	return nil
}

// 根据userid查找购物车
func (r *Repository) FindOrCreate(userid uint) (*Cart, error) {
	var cart *Cart
	//没有则创建
	result := r.DB.Where(&Cart{UserID: userid}).Attrs(NewCart(userid)).FirstOrCreate(&cart)
	if result.Error != nil {
		return nil, result.Error
	}
	return cart, nil
}

type ItemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		DB: db,
	}
}

func (r *ItemRepository) Migration() {
	err := r.DB.AutoMigrate(&Item{})
	if err != nil {
		log.Println(err)
	}
}

// 更新item
func (r *ItemRepository) Update(item *Item) error {
	err := r.DB.Save(item).Error
	if err != nil {
		return err
	}
	return nil
}

// 根据商品SKU和购物车id查找item
func (r *ItemRepository) FindByID(productSKU string, cartID uint) (*Item, error) {
	var item *Item
	err := r.DB.Where(&Item{ProductSKU: productSKU, CartID: cartID}).First(&item).Error
	if err != nil {
		return nil, errors.New("cart item not found")
	}
	return item, nil
}

// 创建item
func (r *ItemRepository) Create(item *Item) error {
	result := r.DB.Create(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 返回购物车中所有item
func (r *ItemRepository) GetItems(cartID uint) ([]Item, error) {
	var items []Item
	//购物车中item不需要分页查询，需要全部显示
	result := r.DB.Where(&Item{CartID: cartID}).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	// //查询购物车检查其中商品是否删除
	// for i, item := range items {
	// 	//这里直接调用gorm的关联功能，减少代码量且可以实时反应数据库中的数据变化，但耦合性高，使用product的repository则增加代码量，解耦合，增加函数调用开销
	// 	//r.DB.Model(item).Association("Product").Find(&items[i].Product)
	// 	r.DB.Model(item).Find(&items[i])

	// }
	return items, nil

}
