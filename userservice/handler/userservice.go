package handler

import (
	"context"
	"os"
	"strconv"
	"time"
	"userservice/config"
	user "userservice/domin"
	pb "userservice/proto"
	"userservice/utils/hash"
	jwt_helper "userservice/utils/jwt"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/protobuf/proto"
)

// 配置文件全局对象
var Appconfig = &config.Configuration{}

// 用户service结构体
type UserService struct {
	repo user.Repository
}

// 实例化service
func NewUserService(r user.Repository) *UserService {
	r.Migration()
	r.InsertSampleData()
	return &UserService{
		repo: r,
	}
}

// 注册
func (s *UserService) Register(ctx context.Context, in *pb.RegisterRequest) (out *pb.RegisterResponse, err error) {
	username := in.Username
	password := in.Password
	out = new(pb.RegisterResponse)
	// 无效用户名
	if user.ValidateUserName(username) {
		out.Status = pb.RegisterResponse_INVALID_USERNAME
		return out, user.ErrInvalidUsername
	}
	// 无效密码
	if user.ValidatePassword(password) {
		out.Status = pb.RegisterResponse_INVALID_PASSWORD
		return out, user.ErrInvalidPassword
	}
	// 用户名存在
	_, err = s.repo.GetByName(username)
	if err == nil {
		out.Status = pb.RegisterResponse_USERNAME_ALREADY_EXISTS
		return out, user.ErrUserExistWithName
	}
	//序列化个人信息
	userInfo, err := proto.Marshal(&pb.UserInfo{Email: in.Email})
	if err != nil {
		return out, err
	}
	// 创建用户
	err = s.repo.Create(user.NewUser(username, password, userInfo))
	if err != nil {
		return out, err
	}
	return out, nil

}

// 登录
func (s *UserService) Login(ctx context.Context, in *pb.LoginRequest) (out *pb.LoginResponse, err error) {
	username := in.Username
	password := in.Password
	usr, err := s.repo.GetByName(username)

	out = new(pb.LoginResponse)
	if err != nil {
		out.Status = pb.LoginResponse_INVALID_USERNAME
		return out, user.ErrUserNotFound
	}
	match := hash.CheckPasswordHash(password+usr.Salt, usr.Password)
	if !match {
		out.Status = pb.LoginResponse_INVALID_PASSWORD
		return out, user.ErrInvalidPassword
	}
	userInfo := &pb.UserInfo{}
	err = proto.Unmarshal(usr.UserInfo, userInfo)
	if err != nil {
		return out, err
	}
	userInfo.UserId = strconv.Itoa(int(usr.ID))
	userInfo.Username = usr.Username

	decodedClaims := jwt_helper.VerifyToken(usr.Token, Appconfig.JWTSettings.SecretKey)
	//token不存在则创建
	if decodedClaims == nil {
		token := jwt_helper.GenerateToken(jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": usr.Username,
			"userId":   strconv.FormatInt(int64(usr.ID), 10), //将整数值转换为十进制表示的字符串，范围为2，36超出则返回空字符串
			"exp":      time.Now().Add(time.Duration(24) * time.Hour).Unix(),
			"iat":      time.Now().Unix(),
			"iss":      os.Getenv("env"), //获取环境变量env的值
			"isAdmin":  usr.IsAdmin,
		}), Appconfig.JWTSettings.SecretKey)
		usr.Token = token
		err := s.repo.Update(&usr)
		if err != nil {
			return out, err
		}
	}
	out.IsAdmin = usr.IsAdmin
	out.Token = usr.Token
	out.Status = pb.LoginResponse_SUCCESS
	out.UserInfo = userInfo
	return out, nil
}
func (s *UserService) GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest) (out *pb.GetUserInfoResponse, err error) {
	uid, _ := strconv.Atoi(in.UserId)
	usr, err := s.repo.GetByID(uint(uid))
	if err != nil {
		return out, err
	}
	userInfo := &pb.UserInfo{}
	err = proto.Unmarshal(usr.UserInfo, userInfo)
	if err != nil {
		return out, err
	}
	out = new(pb.GetUserInfoResponse)
	out.UserInfo = userInfo
	return out, nil
}
