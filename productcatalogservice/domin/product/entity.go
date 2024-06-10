package product

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"productcatalogservice/domin/category"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string            `json:"name"`
	SKU        string            `json:"sku"`
	Desc       string            `json:"desc"`
	StockCount int               `json:"stockCount"`
	PriceUsd   Money             `gorm:"type:varchar(255)" json:"priceUsd"`
	Img        string            `json:"img"`
	CategoryID uint              `json:"category_id"`
	Category   category.Category `json:"-"`
	IsDeleted  bool              `json:"is_deleted"`
}

type Money struct {
	Units int64
	Nanos int32
}

// Scan 方法实现 Scanner 接口
func (m *Money) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan: unable to convert interface to []byte")
	}

	n, err := fmt.Sscanf(string(bytes), "%d.%d", &m.Units, &m.Nanos)
	if err != nil {
		return err
	}

	if n != 2 {
		return errors.New("Scan: invalid money format")
	}

	return nil
}

// Value 方法实现 Valuer 接口
func (m Money) Value() (driver.Value, error) {
	return fmt.Sprintf("%d.%d", m.Units, m.Nanos), nil
}

//	func NewProduct(name, desc string, stockCount int, priceUsd Money, img string, categories []category.Category, pid uint) *Product {
//		return &Product{
//			Name:          name,
//			Desc:          desc,
//			StockCount:    stockCount,
//			PriceUsd:      priceUsd,
//			Img:           img,
//			Categories:    categories,
//			IsDeleted:     false,
//			ProductInfoID: pid,
//		}
//	}
func NewProduct(name, desc string, stockCount int, priceUsd Money, img string) *Product {
	return &Product{
		Name:       name,
		Desc:       desc,
		StockCount: stockCount,
		PriceUsd:   priceUsd,
		Img:        img,
		IsDeleted:  false,
	}
}
