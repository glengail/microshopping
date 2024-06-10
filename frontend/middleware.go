package main

import (
	"fmt"
	jwt_helper "frontend/utils/jwt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type ctxKeyRequestID struct{}

// 定义会话存储
var store = cookie.NewStore([]byte("secret"))

// 令牌桶限流
type TokenBucket struct {
	cap             int64         //令牌桶容量
	fillInterval    time.Duration //填充令牌间隔时间
	avaliableTokens int64         //当前令牌数量
	lastfilltime    time.Time     //最后取令牌时间
	mu              sync.Mutex    //互斥锁
}

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
		//token := c.GetHeader("Authorization")
		session := sessions.Default(c)
		v := session.Get("token")
		token := ""
		if v != nil {
			token = v.(string)
		}
		if token != "" {
			decodedClaim := jwt_helper.VerifyToken(token, secretKey)
			if decodedClaim != nil {
				c.Set(userIdText, decodedClaim.UserId)
				c.Next()
				return
			}
			//该方法包含c.Abort()中断请求
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "你没有权限访问",
			})
			c.Abort()
			return

		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "请先登录！",
			})
			c.Abort()
			return
		}

	}
}

func NewTokenBucket(cap int64, fillInterval time.Duration) *TokenBucket {
	return &TokenBucket{
		cap:             cap,
		fillInterval:    fillInterval,
		avaliableTokens: cap,
		lastfilltime:    time.Now(),
	}
}

// 尝试获取一个令牌，如果令牌桶中有可用令牌则返回true，否则返回false
func (tb *TokenBucket) Take() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	num := tb.tokensToAdd(time.Now())
	tb.avaliableTokens += num
	if tb.avaliableTokens > tb.cap {
		tb.avaliableTokens = tb.cap
	}
	if tb.avaliableTokens > 0 {
		tb.avaliableTokens--
		return true
	}
	return false
}

// 计算上次填充至当前时间间隔内应该填充的令牌数量
func (tb *TokenBucket) tokensToAdd(now time.Time) int64 {
	d := now.Sub(tb.lastfilltime)
	if d < 0 {
		return 0
	}
	num := int64(d / tb.fillInterval)
	tb.lastfilltime = tb.lastfilltime.Add(time.Duration(num) * tb.fillInterval)
	return num

}

func RateLimitMiddleWare(tb *TokenBucket) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !tb.Take() {
			ctx.JSON(302, gin.H{
				"msg": fmt.Sprintln("访问次数频繁"),
			})
			//Abort() 方法用于终止当前请求并阻止后续的中间件函数或处理程序执行
			ctx.Abort()
			return
		}
		//Next() 方法通常用于在中间件函数中将控制权交给链中的下一个中间件或处理程序，如果后续没有中间件，则处理函数执行
		ctx.Next()
	}
}
