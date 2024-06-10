package handler

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/rand"
	pb "recommendationservice/proto"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

type RecommendationService struct {
	ProductCatalogService pb.ProductCatalogServiceClient
}

func (c *RecommendationService) ListRecommendations(ctx context.Context, in *pb.ListRecommendationsRequest) (out *pb.ListRecommendationsResponse, err error) {
	maxRecommendationCount := 5
	out = new(pb.ListRecommendationsResponse)
	//查询商品类别
	// p, err := c.ProductCatalogService.GetProduct(ctx, &pb.GetProductRequest{
	// 	Id: in.ProductIds[0],
	// })
	// fmt.Printf("p: %v\n", p)
	ids := in.ProductIds
	if ids == nil {
		ids = make([]string, 10)
	}
	catalog, err := c.ProductCatalogService.ListProducts(ctx, &pb.PageRequest{
		Page:     -1,
		PageSize: int32(len(ids)),
	})
	if err != nil {
		return nil, err
	}
	//fmt.Printf("catalog.Products: %v\n", catalog.Page)
	filteredProductIds := make([]string, 0, len(catalog.GetPage().GetProducts()))
	for _, p := range catalog.GetPage().GetProducts() {
		//去掉已有，添加相关推荐
		if contains(in.ProductIds, p.Id) {
			fmt.Printf("p.Id: %v\n", p.Id)
			continue
		}
		filteredProductIds = append(filteredProductIds, p.Id)
	}
	productIds := sample(filteredProductIds, maxRecommendationCount)
	logger.Printf("[recv ListRecommendations] product_ids: %v", productIds)
	out.ProductIds = productIds
	return out, nil
}

// 判断是否包含
func contains(source []string, target string) bool {
	for _, v := range source {
		if v == target {
			return true
		}
	}
	return false
}

// 示例
func sample(source []string, c int) []string {
	n := len(source)
	if c >= n {
		return source
	}
	indicts := make([]int, n)
	for i := 0; i < n; i++ {
		indicts[i] = i
	}
	//fisher yats洗牌算法，每个下标被抽中与当前下标交换的概率相同，也不会重复
	for i := n - 1; i > 0; i-- {
		//随机一个数0到当前的下标与当前下标进行交换
		j := rand.Intn(i + 1)
		indicts[i], indicts[j] = indicts[j], indicts[i]
	}
	result := make([]string, 0, c)
	for i := 0; i < c; i++ {
		result = append(result, source[indicts[i]])
	}
	return result

}
