package main

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
	"time"

	"frontend/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "frontend/proto"
)

const (
	name    = "frontend"
	version = "1.0.0"

	defaultCurrency = "USD"
	cookieMaxAge    = 60 * 60 * 48

	cookiePrefix    = "shop_"
	cookieSessionID = cookiePrefix + "session-id"
	cookieCurrency  = cookiePrefix + "currency"
)

var (
	whitelistedCurrencies = map[string]bool{
		"USD": true,
		"EUR": true,
		"CAD": true,
		"JPY": true,
		"GBP": true,
		"TRY": true,
	}
	Appconfig = &config.Configuration{}
)

type ctxKeySessionID struct {
	sessionID string
}

// 前端server
type FrontendServer struct {
	adService             pb.AdServiceClient
	cartService           pb.CartServiceClient
	checkoutService       pb.CheckoutServiceClient
	currencyService       pb.CurrencyServiceClient
	productCatalogService pb.ProductCatalogServiceClient
	recommendationService pb.RecommendationServiceClient
	shippingService       pb.ShippingServiceClient
	userService           pb.UserServiceClient
	AppConfig             *config.Configuration
}

// 获得grpc连接
func GetGrpcConn(consulClient *api.Client, serviceName string, serviceTag string) *grpc.ClientConn {
	service, _, err_service := consulClient.Health().Service(serviceName, serviceTag, true, nil)
	if err_service != nil {
		fmt.Println("获取健康服务报错：", err_service)
		return nil
	}
	// fmt.Println(service[0].Service)
	s := service[0].Service
	address := s.Address + ":" + strconv.Itoa(s.Port)
	fmt.Printf("address: %v\n", address)
	//链接grpc服务
	grpcConn, _ := grpc.Dial(address, grpc.WithInsecure())

	return grpcConn
}

func main() {

	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout

	//1.初始化consul配置
	consulConfig := api.DefaultConfig()
	//2.创建consul对象
	consulClient, err_consul := api.NewClient(consulConfig)
	if err_consul != nil {
		fmt.Println("consul创建对象报错：", err_consul)
		return
	}

	svc := &FrontendServer{
		adService:             pb.NewAdServiceClient(GetGrpcConn(consulClient, "adservice", "adservice")),
		cartService:           pb.NewCartServiceClient(GetGrpcConn(consulClient, "cartservice", "cartservice")),
		checkoutService:       pb.NewCheckoutServiceClient(GetGrpcConn(consulClient, "checkoutservice", "checkoutservice")),
		currencyService:       pb.NewCurrencyServiceClient(GetGrpcConn(consulClient, "currencyservice", "currencyservice")),
		productCatalogService: pb.NewProductCatalogServiceClient(GetGrpcConn(consulClient, "productcatalogservice", "productcatalogservice")),
		recommendationService: pb.NewRecommendationServiceClient(GetGrpcConn(consulClient, "recommendationservice", "recommendationservice")),
		shippingService:       pb.NewShippingServiceClient(GetGrpcConn(consulClient, "shippingservice", "shippingservice")),
		userService:           pb.NewUserServiceClient(GetGrpcConn(consulClient, "userservice", "userservice")),
	}

	r := gin.Default()

	r.FuncMap = template.FuncMap{
		"renderMoney":        renderMoney,
		"renderCurrencyLogo": renderCurrencyLogo,
	}

	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")
	// 首页
	r.GET("/", svc.HomeHandler)
	// 商品
	r.GET("/product/:id", svc.GetProductHandler)
	//商品搜索
	r.GET("/product1/search", svc.SearchProductsHandler)
	// 查看购物车
	r.GET("/cart", AuthUserMiddleWare(Appconfig.JWTSettings.SecretKey), svc.viewCartHandler)
	// 添加购物车
	r.POST("/cart", AuthUserMiddleWare(Appconfig.JWTSettings.SecretKey), svc.addToCartHandler)
	// 清空购物车
	r.POST("/cart/empty", AuthUserMiddleWare(Appconfig.JWTSettings.SecretKey), svc.emptyCartHandler)
	// 设置货币种类
	r.POST("/setCurrency", svc.setCurrencyHandler)
	// // 退出登录
	r.GET("/logout", svc.logoutHandler)
	// // 结账
	r.POST("/cart/checkout", AuthUserMiddleWare(Appconfig.JWTSettings.SecretKey), svc.placeOrderHandler)
	// 登录
	//创建一个容量位20的令牌桶，每秒填充一个
	loginTb := NewTokenBucket(20, time.Second)
	r.POST("/login", RateLimitMiddleWare(loginTb), svc.loginHandler)
	r.GET("/login", svc.Login)
	// 注册
	//创建一个容量位20的令牌桶，每秒填充一个
	registerTb := NewTokenBucket(20, time.Second)
	r.POST("/register", RateLimitMiddleWare(registerTb), svc.registerHandler)
	r.GET("/register", svc.Register)
	// 搜索
	r.GET("search", svc.SearchProductsHandler)
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("gin启动失败: %v", err)
	}

}
