package handler

import (
	cart "cartservice/domin"
	pb "cartservice/proto"
	context "context"
	"fmt"
	"strconv"
)

type CartService struct {
	CartRepo              cart.Repository
	ItemRepo              cart.ItemRepository
	ProductCatalogService pb.ProductCatalogServiceClient
}

func NewCartService(cartRepo cart.Repository, itemRepo cart.ItemRepository, productCatalogService pb.ProductCatalogServiceClient) *CartService {
	cartRepo.Migration()
	itemRepo.Migration()
	return &CartService{
		CartRepo:              cartRepo,
		ItemRepo:              itemRepo,
		ProductCatalogService: productCatalogService,
	}
}

// 添加商品
func (s *CartService) AddItem(ctx context.Context, in *pb.AddItemRequest) (out *pb.Empty, err error) {

	uid, _ := strconv.ParseUint(in.UserId, 10, 64)
	sku := in.Item.ProductSku
	//商品非法
	p, err := s.ProductCatalogService.GetProduct(ctx, &pb.GetProductRequest{Id: sku})
	if err != nil {
		return out, cart.ErrInvalidItems
	}
	//数量非法
	if in.Item.Quantity <= 0 {
		return out, cart.ErrCountInvalid
	}
	fmt.Printf("in.Item: %v\n", in.Item)
	//商品已存在
	item, err := s.ItemRepo.FindByID(p.Id, uint(uid))
	if item != nil && err == nil {
		item.Count += int(in.Item.Quantity)
		s.ItemRepo.Update(item)
		return out, nil
	}

	//添加Item
	userCart, err := s.CartRepo.FindOrCreate(uint(uid))
	if err != nil {
		return out, err
	}
	if err := s.ItemRepo.Create(cart.NewCartItem(sku, userCart.ID, int(in.Item.Quantity))); err != nil {
		return out, err
	}
	return out, nil
}

// 清空购物车
func (s *CartService) EmptyCart(ctx context.Context, in *pb.EmptyCartRequest) (out *pb.Empty, err error) {
	uid, _ := strconv.ParseUint(in.UserId, 10, 64)
	cart, err := s.CartRepo.FindOrCreate(uint(uid))
	if err != nil {
		return out, err
	}
	err = s.CartRepo.Delete(cart)
	if err != nil {
		return out, err
	}
	return out, nil
}

// 获得购物车
func (s *CartService) GetCart(ctx context.Context, in *pb.GetCartRequest) (out *pb.Cart, err error) {
	uid, _ := strconv.ParseUint(in.UserId, 10, 64)
	cart, err := s.CartRepo.FindOrCreate(uint(uid))
	if err != nil {
		return out, err
	}
	items, err := s.ItemRepo.GetItems(cart.ID)
	if err != nil {
		return out, err
	}
	results := cartItemsToPb(items)
	out = new(pb.Cart)
	out.UserId = in.UserId
	out.Items = results
	return out, nil
}

func cartItemsToPb(items []cart.Item) []*pb.CartItem {
	cartItems := make([]*pb.CartItem, 0, len(items))
	for _, item := range items {
		cartItems = append(cartItems, cartItemToPb(item))
	}
	return cartItems
}

// cartItem转pb.cartItem
func cartItemToPb(item cart.Item) *pb.CartItem {
	return &pb.CartItem{
		ProductSku: item.ProductSKU,
		Quantity:   int32(item.Count),
	}
}

// pb.cartItem转cartItem
func pbToCartItem(item *pb.CartItem) cart.Item {
	if item != nil {
		return cart.Item{
			ProductSKU: item.ProductSku,
			Count:      int(item.Quantity),
		}
	}
	return cart.Item{}
}
