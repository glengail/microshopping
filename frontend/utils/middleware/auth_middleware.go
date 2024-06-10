package middleware

import (
	jwt_helper "frontend/utils/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthAdminMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token != "" {
			decodedClaim := jwt_helper.VerifyToken(token, secretKey)
			if decodedClaim != nil && decodedClaim.IsAdmin {
				c.Next()
				c.Abort()
				return
			}
			//该方法包含c.Abort()中断请求
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "你没有权限访问",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "未授权！",
		})
	}
}
func AuthUserMiddleWare(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token != "" {
			decodedClaim := jwt_helper.VerifyToken(token, secretKey)
			if decodedClaim != nil {
				c.Set("userId", decodedClaim.UserId)
				c.Next()
				c.Abort()
				return
			}
			//该方法包含c.Abort()中断请求
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "你没有权限访问",
			})
			return

		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "未授权！",
		})
	}
}
