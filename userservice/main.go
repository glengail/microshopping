package main

import (
	"fmt"
	"net"
	"strconv"
	"userservice/config"
	user "userservice/domin"
	handler "userservice/handler"
	pb "userservice/proto"
	databasehandler "userservice/utils/database_handler"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const PORT = 50021
const ADDRESS = "127.0.0.1"

// --------------连接数据库----------------
type Databases struct {
	userRepo *user.Repository
}

// 实例化
func CreateDB() *Databases {
	cfgFile := "config/config.yaml"
	appcfg, err := config.GetAllConfigValues(cfgFile)
	if err != nil {
		panic(err)
	}
	handler.Appconfig = appcfg
	db := databasehandler.NewMysqlDB(appcfg.DatabaseSettings.DatabaseURI)
	return &Databases{
		userRepo: user.NewUserRepository(db),
	}
}
func main() {
	db := CreateDB()
	userService := handler.NewUserService(*db.userRepo)
	ipport := ADDRESS + ":" + strconv.Itoa(PORT)

	// 注册到consul上
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
		Tags:    []string{"userservice"},
		Name:    "userservice",
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
	pb.RegisterUserServiceServer(grpcServer, userService)

	// 设置监听
	listien, err := net.Listen("tcp", ipport)
	if err != nil {
		fmt.Println("监听报错:", err)
		return
	}
	defer listien.Close()

	// 启动服务
	fmt.Println("服务启动成功...")

	err_grpc := grpcServer.Serve(listien)
	if err_grpc != nil {
		fmt.Println("grpc服务启动报错:", err)
		return
	}
}
