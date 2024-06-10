package category

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

//生成表

func (r *Repository) Migration() {
	err := r.DB.AutoMigrate(&Category{})
	if err != nil {
		log.Println(err)
	}
	r.InsertSampleData()
}

func (r *Repository) Create(c *Category) error {
	result := r.DB.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 插入测试分类
func (r *Repository) InsertSampleData() {
	categories := []Category{
		{Name: "衣服", Desc: "clothing"},
		{Name: "家具", Desc: "jiaju"},
	}
	for _, c := range categories {
		r.DB.Where(Category{Name: c.Name}).Attrs(Category{Name: c.Name, Desc: c.Desc}).FirstOrCreate(&c)
	}
}

//通过名称查询分类

func (r *Repository) GetByName(name string) []Category {
	var c []Category
	r.DB.Where("name =?", name).Find(&c)
	return c
}

//批量创建商品分类

func (r *Repository) BulkCreate(categories []*Category) (int, error) {
	var count int64
	err := r.DB.Create(categories).Count(&count).Error
	return int(count), err
}

// 获得分页商品分类
func (r *Repository) GetAll(pageIndex, pageSize int) ([]Category, int) {
	var categories []Category
	var count int64

	r.DB.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Count(&count)
	return categories, int(count)
}
