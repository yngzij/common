package jwt

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTConfig struct {
	ClaimsKey string
	ExpiresAt int64
	Secret    string
}

type SessionJWTClaims struct {
	UserID             string `json:"user_id"`
	OpenID             string `json:"open_id"`
	ExpiresAt          int64  `json:"expires_at"`
	jwt.StandardClaims `json:"claims"`
}

func GenerateJWT(key string, claims SessionJWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func Claims(ctx *gin.Context) (*SessionJWTClaims, error) {
	var claims *SessionJWTClaims
	if val, ok := ctx.Get("user"); ok {
		claims = val.(*SessionJWTClaims)
	} else {
		return nil, fmt.Errorf("claims not find")
	}
	return claims, nil
}

func JWT(config JWTConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtToken := ctx.Request.Header.Get("token")
		if len(jwtToken) == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": 205,
				"msg":  "Not Token",
			})
			return
		}
		tokenString := strings.TrimSpace(jwtToken)
		sessionsClaims := &SessionJWTClaims{}
		_, err := jwt.ParseWithClaims(tokenString, sessionsClaims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(config.Secret), nil
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": 207,
				"msg":  "parse JTW error",
			})
			return nil, nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": 206,
				"msg":  fmt.Sprintf("Parse Claims Error:%s", err),
			})
			return
		}
		if sessionsClaims.ExpiresAt > time.Now().Unix() {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": 208,
				"msg":  fmt.Sprintf("Token already expired"),
			})
			return
		}
		ctx.Set(config.ClaimsKey, sessionsClaims)
	}
}
