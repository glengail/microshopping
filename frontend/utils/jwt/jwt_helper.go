package jwt_helper

import (
	"encoding/json"
	"log"

	"github.com/dgrijalva/jwt-go"
)

type DecodedToken struct {
	Iat      int    `json:"iat"` //签发jwt时间
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Iss      string `json:"iss"` //签发人
	IsAdmin  bool   `json:"isAdmin"`
}

func GenerateToken(claims *jwt.Token, secret string) (token string) {
	//将secret转换为byte数组
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	//生成token
	token, _ = claims.SignedString(hmacSecret)
	return
}

// 解析token，对比原码token和secret
func VerifyToken(token string, secret string) *DecodedToken {
	//将secret转换为byte数组
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	decoded, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		//应该返回hamcSecret
		return hmacSecret, nil
	})
	if err != nil {
		return nil
	}
	if !decoded.Valid {
		return nil
	}
	var decodedToken *DecodedToken
	//将解析的声明转换为结构体jwt.MapClaims
	//jwt.MapClaims实际是map[]interface{}键值对
	decodedClaims, _ := decoded.Claims.(jwt.MapClaims)
	jsonStr, _ := json.Marshal(decodedClaims)
	//解析成具体结构体看是否符合decodedToken
	jsonErr := json.Unmarshal(jsonStr, &decodedToken)
	if jsonErr != nil {
		log.Println(jsonErr)
		return nil
	}
	return decodedToken
}
