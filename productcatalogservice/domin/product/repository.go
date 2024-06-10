package product

import (
	"log"

	"gorm.io/gorm"
)

var isInitialized bool // 包级别的变量，用于标记函数是否已经执行过
type Repository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

// 生成表
func (r *Repository) Migration() {
	err := r.DB.AutoMigrate(&Product{
		PriceUsd: Money{},
	})
	//r.InsertSampleData()
	if err != nil {
		log.Println(err)
	}
}

// 插入测试分类
func (r *Repository) InsertSampleData(products []*Product) {

	// products := []Product{
	// 	{
	// 		SKU:  "OLJCESPC7Z",
	// 		Name: "太阳镜",
	// 		Desc: "这款时尚的飞行员太阳镜为您的服装增添现代感。",
	// 		Img:  "/static/img/products/sunglasses.jpg",
	// 		PriceUsd: &Money{
	// 			Money: pb.Money{
	// 				CurrencyCode: "USD",
	// 				Units:        19,
	// 				Nanos:        990000000,
	// 			},
	// 		},
	// 		CategoryID: 1,
	// 	},
	// 	{
	// 		SKU:  "66VCHSJNUP",
	// 		Name: "背心",
	// 		Desc: "完美剪裁的棉质背心，鸡心领。",
	// 		Img:  "/static/img/products/tank-top.jpg",
	// 		PriceUsd: &Money{
	// 			Money: pb.Money{
	// 				CurrencyCode: "USD",
	// 				Units:        18,
	// 				Nanos:        990000000,
	// 			},
	// 		},
	// 		CategoryID: 2,
	// 	},
	// }
	if !isInitialized {
		for _, p := range products {
			if p != nil {
				r.DB.Where(Product{Name: p.Name}).Attrs(Product{Name: p.Name, Desc: p.Desc, Img: p.Img, PriceUsd: p.PriceUsd, CategoryID: 1}).FirstOrCreate(&p)
			}
		}
		isInitialized = true
	}
}

// 创建产品

func (r *Repository) Create(p *Product) error {
	result := r.DB.Create(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 更新
func (r *Repository) Update(p *Product) error {
	savedProduct, err := r.FindBySku(p.SKU)
	if err != nil {
		return err
	}
	return r.DB.Model(savedProduct).Updates(p).Error

}

// 返回搜索结果
func (r *Repository) SearchByString(str string, pageIndex, pageSize int) ([]Product, int64) {
	var products []Product
	var count int64
	convertStr := "%" + str + "%"
	r.DB.Where("IsDeleted = ?", 0).Where("name LIKE ? OR Sku LIKE ?", convertStr, convertStr).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)
	return products, count
}

// 查询所有商品
func (r *Repository) GetAll(pageIndex, pageSize int) ([]*Product, int64) {
	var products []*Product
	var count int64
	r.DB.Where("IsDeleted = ?", 0).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)
	return products, count
}

// 根据SKU删除
func (r *Repository) DeleteBySku(sku string) error {
	savedProduct, err := r.FindBySku(sku)
	if err != nil {
		return err
	}
	savedProduct.IsDeleted = true
	return r.DB.Save(savedProduct).Error
}

// 根据SKU查找
func (r *Repository) FindBySku(sku string) (*Product, error) {
	var p *Product
	err := r.DB.Where("IsDeleted = ?", 0).Where(&Product{SKU: sku}).First(&p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}
