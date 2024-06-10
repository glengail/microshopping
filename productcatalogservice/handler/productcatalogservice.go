package handler

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	product "productcatalogservice/domin/product"
	pb "productcatalogservice/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

var reloadCatalog bool

// 日志
var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

const (
	SIGUSR1 = syscall.Signal(0xa)
	SIGUSR2 = syscall.Signal(0xc)
)

// 商品分类结构体
type ProductCatalogService struct {
	// categoryRepo category.Repository
	productRepo product.Repository
	// productService product.Service
	sync.Mutex
	products []*pb.Product
}

func NewProductCatalogService(productRepo product.Repository) *ProductCatalogService {
	productRepo.Migration()
	return &ProductCatalogService{
		productRepo: productRepo,
	}
}

// 商品列表
func (s *ProductCatalogService) ListProducts(ctx context.Context, in *pb.PageRequest) (out *pb.ListProductsResponse, e error) {
	out = new(pb.ListProductsResponse)
	out.Page = new(pb.PageResponse)
	data := s.pb_productsToProducts(s.parseCatalog())
	s.productRepo.InsertSampleData(data)
	products, count := s.productRepo.GetAll(int(in.Page), int(in.PageSize))
	out.Page.Products = s.productsTopbProducts(products)
	for _, p := range out.Page.Products {
		fmt.Printf("p: %v\n", p)
	}
	out.Page.Total = count
	return out, nil
}

// 获得单个商品
func (s *ProductCatalogService) GetProduct1(ctx context.Context, in *pb.GetProductRequest) (out *pb.Product, e error) {
	var found *pb.Product
	out = new(pb.Product)
	products := s.parseCatalog()
	for _, p := range products {
		if in.Id == p.Id {
			found = p
		}
	}
	if found == nil {
		return out, status.Errorf(codes.NotFound, "no product with ID %s", in.Id)
	}
	out.Id = found.Id
	out.Name = found.Name
	out.Categories = found.Categories
	out.Description = found.Description
	out.Picture = found.Picture
	out.PriceUsd = found.PriceUsd
	return out, nil
}

// 获得单个商品
func (s *ProductCatalogService) GetProduct(ctx context.Context, in *pb.GetProductRequest) (out *pb.Product, e error) {
	out = new(pb.Product)
	found, err := s.productRepo.FindBySku(in.Id)
	if err != nil {
		return out, status.Errorf(codes.NotFound, "商品不存在 %s", in.Id)
	}
	out.Id = found.SKU
	out.Name = found.Name
	//out.Categories = found.
	out.Description = found.Desc
	out.Picture = found.Img
	out.PriceUsd = s.MoneyToPbMoney(found.PriceUsd)
	return out, nil
}

// 搜索商品
func (s *ProductCatalogService) SearchProducts(ctx context.Context, in *pb.SearchProductsRequest) (out *pb.SearchProductsResponse, e error) {
	// var ps []*pb.Product
	// products := s.parseCatalog()
	// for _, p := range products {
	// 	if strings.Contains(strings.ToLower(p.Name), strings.ToLower(in.Query)) ||
	// 		strings.Contains(strings.ToLower(p.Description), strings.ToLower(in.Query)) {
	// 		ps = append(ps, p)
	// 	}
	// }
	// out.Results = ps
	// return out, nil

	products, count := s.productRepo.SearchByString(in.Query, int(in.Page.Page), int(in.Page.PageSize))
	results := make([]*pb.Product, len(products))
	for _, p := range products {
		// categories := make([]string, len(p.Categories))
		// for _, c := range p.Categories {
		// 	categories = append(categories, c.Name)
		// }
		results = append(results, &pb.Product{
			Id:          p.SKU,
			Name:        p.Name,
			Description: p.Desc,
			Picture:     p.Img,
			PriceUsd: &pb.Money{
				CurrencyCode: "USD",
				Units:        p.PriceUsd.Units,
				Nanos:        p.PriceUsd.Nanos,
			},
		})
	}
	out = new(pb.SearchProductsResponse)
	out.Page = &pb.PageResponse{
		Total:    count,
		Products: results,
	}

	// results := s.productService.SearchProduct(in.Query, &pagination.Pages{Page: int(in.Page.Page), PageSize: int(in.Page.PageSize)})
	// data, err := interfaceToAny(results.Items.([]product.Product))
	// if err != nil {
	// 	return nil, err
	// }
	// out = new(pb.SearchProductsResponse)
	// out.Page = &pb.PageResponse{
	// 	Total: int64(results.TotalCount),
	// 	Data:  data,
	// }
	return out, nil
}

// product转pb_product
func (s *ProductCatalogService) productTopbProduct(p *product.Product) *pb.Product {
	if p != nil {
		return &pb.Product{
			Id:          p.SKU,
			Name:        p.Name,
			Description: p.Desc,
			Picture:     p.Img,
			PriceUsd:    s.MoneyToPbMoney(p.PriceUsd),
			Categories:  []string{p.Category.Name},
		}
	}
	return &pb.Product{}
}

// products转pb_products
func (s *ProductCatalogService) productsTopbProducts(ps []*product.Product) []*pb.Product {
	products := make([]*pb.Product, 0, len(ps))
	for _, p := range ps {
		products = append(products, s.productTopbProduct(p))
	}
	return products
}

// pb_product转product
func (s *ProductCatalogService) pb_productToProduct(p *pb.Product) *product.Product {
	if p != nil {
		return product.NewProduct(p.Name, p.Description, 0, s.PbMoneyToMoney(p.PriceUsd), p.Picture)
	}
	return &product.Product{}

}
func (s *ProductCatalogService) PbMoneyToMoney(money *pb.Money) product.Money {
	if money != nil {
		return product.Money{
			Units: money.Units,
			Nanos: money.Nanos,
		}
	}
	return product.Money{
		Units: 1,
		Nanos: 0,
	}

}

// money转pbMoney
func (s *ProductCatalogService) MoneyToPbMoney(money product.Money) *pb.Money {
	return &pb.Money{
		CurrencyCode: "USD",
		Units:        money.Units,
		Nanos:        money.Nanos,
	}
}

//pb_products转products

func (s *ProductCatalogService) pb_productsToProducts(ps []*pb.Product) []*product.Product {
	products := make([]*product.Product, 0, len(ps))
	for _, p := range ps {
		if p != nil {
			products = append(products, s.pb_productToProduct(p))
		}
	}
	return products
}

// 读配置文件
func (s *ProductCatalogService) readCatalogFile() (*pb.ListProductsResponse, error) {
	s.Lock()
	defer s.Unlock()
	catalogJSON, err := os.ReadFile("data/products.json")
	if err != nil {
		logger.Printf("打开商品 json 文件失败: %v", err)
		return nil, err
	}
	catalog := &pb.ListProductsResponse{Page: &pb.PageResponse{}}
	if err := protojson.Unmarshal(catalogJSON, catalog.Page); err != nil {
		logger.Printf("解析商品 JSON 文件失败: %v", err)
		return nil, err
	}
	logger.Printf("解析商品 JSON 文件成功")
	//s.productRepo.InsertSampleData(catalog.Page.Products)
	return catalog, nil
}

// 解析配置文件
func (s *ProductCatalogService) parseCatalog() []*pb.Product {
	if reloadCatalog || len(s.products) == 0 {
		catalog, err := s.readCatalogFile()
		if err != nil {
			return []*pb.Product{}
		}
		s.products = catalog.Page.Products
	}
	return s.products
}

// 初始化
func init() {
	signs := make(chan os.Signal, 1)
	signal.Notify(signs, SIGUSR1, SIGUSR2)
	go func() {
		for {
			sig := <-signs
			switch sig {
			case SIGUSR1:
				logger.Println("可以加载商品信息")
				reloadCatalog = true
			case SIGUSR2:
				logger.Println("不能加载商品信息")
				reloadCatalog = false
			}
		}
	}()

}
