package main

import (
	"fmt"
	"net"
	"productcatalogservice/config"
	"productcatalogservice/domin/category"
	"productcatalogservice/domin/product"
	"productcatalogservice/handler"
	pb "productcatalogservice/proto"
	databasehandler "productcatalogservice/utils/database_handler"
	"strconv"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const PORT = 50015
const ADDRESS = "127.0.0.1"

// --------------连接数据库----------------
type Databases struct {
	categoryRepo *category.Repository
	productRepo  *product.Repository
}

// 配置文件全局对象
var Appconfig = &config.Configuration{}

// 实例化
func CreateDB() *Databases {
	cfgFile := "config/config.yaml"
	appcfg, err := config.GetAllConfigValues(cfgFile)
	if err != nil {
		panic(err)
	}
	Appconfig = appcfg
	db := databasehandler.NewMysqlDB(appcfg.DatabaseSettings.DatabaseURI)
	return &Databases{
		categoryRepo: category.NewRepository(db),
		productRepo:  product.NewProductRepository(db),
	}
}
func main() {
	db := CreateDB()
	productService := handler.NewProductCatalogService(*db.productRepo)

	ipport := ADDRESS + ":" + strconv.Itoa(PORT)
	// -------------注册到consul上---------------
	// 初始化consul配置
	consulConfig := api.DefaultConfig()

	// 创建consul对象
	consulClient, err_consul := api.NewClient(consulConfig)
	if err_consul != nil {
		fmt.Println("consul创建对象报错：", err_consul)
		return
	}

	// 告诉consul即将注册到服务到信息
	reg := api.AgentServiceRegistration{
		Tags:    []string{"productcatalogservice"},
		Name:    "productcatalogservice",
		Address: ADDRESS,
		Port:    PORT,
	}

	// 注册grpc服务到consul上
	err_agent := consulClient.Agent().ServiceRegister(&reg)
	if err_agent != nil {
		fmt.Println("consul注册grpc失败：", err_agent)
		return
	}

	//-----------------------grpc代码----------------------------------
	// 初始化grpc对象
	grpcServer := grpc.NewServer()

	// 注册服务

	pb.RegisterProductCatalogServiceServer(grpcServer, productService)
	// 设置监听
	listien, err := net.Listen("tcp", ipport)
	if err != nil {
		fmt.Println("监听报错:", err)
		return
	}
	defer listien.Close()

	// 启动服务
	fmt.Println("服务启动成功。。。")

	err_grpc := grpcServer.Serve(listien)
	if err_grpc != nil {
		fmt.Println("grpc服务启动报错:", err)
		return
	}
}
