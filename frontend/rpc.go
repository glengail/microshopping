package main

import (
	"context"
	"fmt"
	pb "frontend/proto"
	"time"

	"github.com/pkg/errors"
	_ "google.golang.org/protobuf/proto"
)

const (
	//避免重复转换，如果考虑时间汇率变化需要为false
	avoidNoopCurrencyConversionRPC = false
)

func (fe *FrontendServer) getCurrencies(ctx context.Context) ([]string, error) {
	currs, err := fe.currencyService.GetSupportedCurrencies(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}
	var out []string
	for _, c := range currs.CurrencyCodes {
		if _, ok := whitelistedCurrencies[c]; ok {
			out = append(out, c)
		}
	}
	return out, nil
}

func (fe *FrontendServer) getProducts(ctx context.Context, page *pb.PageRequest) ([]*pb.Product, error) {
	resp, err := fe.productCatalogService.ListProducts(ctx, page)
	return resp.Page.Products, err
}

func (fe *FrontendServer) getProduct(ctx context.Context, id string) (*pb.Product, error) {
	resp, err := fe.productCatalogService.GetProduct(ctx, &pb.GetProductRequest{Id: id})
	return resp, err
}
func (fe *FrontendServer) searchProducts(ctx context.Context, qt string, page *pb.PageRequest) ([]*pb.Product, error) {
	resp, err := fe.productCatalogService.SearchProducts(ctx, &pb.SearchProductsRequest{Query: qt, Page: page})
	results := make([]*pb.Product, len(resp.GetPage().GetProducts()))
	//分页解析
	for i, p := range resp.GetPage().GetProducts() {
		results[i] = p
	}
	// 创建一个新的 Product 实例
	return results, err
}
func (fe *FrontendServer) getCart(ctx context.Context, userID string) ([]*pb.CartItem, error) {
	resp, err := fe.cartService.GetCart(ctx, &pb.GetCartRequest{UserId: userID})
	return resp.GetItems(), err
}

func (fe *FrontendServer) emptyCart(ctx context.Context, userID string) error {
	_, err := fe.cartService.EmptyCart(ctx, &pb.EmptyCartRequest{UserId: userID})
	return err
}

func (fe *FrontendServer) insertCart(ctx context.Context, userID, productID string, quantity int32) error {
	_, err := fe.cartService.AddItem(ctx, &pb.AddItemRequest{
		UserId: userID,
		Item: &pb.CartItem{
			ProductId: productID,
			Quantity:  quantity,
		},
	})
	return err
}

func (fe *FrontendServer) convertCurrency(ctx context.Context, money *pb.Money, currency string) (*pb.Money, error) {
	if avoidNoopCurrencyConversionRPC && money.GetCurrencyCode() == currency {
		return money, nil
	}
	return fe.currencyService.Convert(ctx, &pb.CurrencyConversionRequest{
		From:   money,
		ToCode: currency,
	})
}

func (fe *FrontendServer) getShippingQuote(ctx context.Context, items []*pb.CartItem, currency string) (*pb.Money, error) {
	quote, err := fe.shippingService.GetQuote(ctx, &pb.GetQuoteRequest{
		Address: nil,
		Items:   items,
	})
	if err != nil {
		return nil, err
	}
	localized, err := fe.convertCurrency(ctx, quote.GetCostUsd(), currency)
	return localized, errors.Wrap(err, "failed to convert currency for shipping cost")
}

func (fe *FrontendServer) getRecommendations(ctx context.Context, userID string, productIDs []string) ([]*pb.Product, error) {
	fmt.Printf("productIDs: %v\n", productIDs)
	resp, err := fe.recommendationService.ListRecommendations(ctx, &pb.ListRecommendationsRequest{
		UserId:     userID,
		ProductIds: productIDs,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*pb.Product, len(resp.GetProductIds()))
	for i, v := range resp.GetProductIds() {
		p, err := fe.getProduct(ctx, v)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get recommended product info (#%s)", v)
		}
		out[i] = p
	}
	if len(out) > 4 {
		out = out[:4] // take only first four to fit the UI
	}
	return out, err
}

func (fe *FrontendServer) getAd(ctx context.Context, ctxKeys []string) ([]*pb.Ad, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	resp, err := fe.adService.GetAds(ctx, &pb.AdRequest{
		ContextKeys: ctxKeys,
	})
	return resp.GetAds(), errors.Wrap(err, "failed to get ads")
}

func (fe *FrontendServer) login(ctx context.Context, username string, password string) (*pb.LoginResponse, error) {
	resp, err := fe.userService.Login(ctx, &pb.LoginRequest{
		Username: username,
		Password: password,
	})
	return resp, err
}
func (fe *FrontendServer) register(ctx context.Context, username string, password string, email string) (*pb.RegisterResponse, error) {
	resp, err := fe.userService.Register(ctx, &pb.RegisterRequest{
		Username: username,
		Password: password,
		Email:    email,
	})
	return resp, err
}
func (fe *FrontendServer) getUserInfo(ctx context.Context, userid string) (*pb.GetUserInfoResponse, error) {
	resp, err := fe.userService.GetUserInfo(ctx, &pb.GetUserInfoRequest{UserId: userid})
	return resp, err
}
