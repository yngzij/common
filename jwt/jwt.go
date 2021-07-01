package jwt

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTConfig struct {
	ClaimsKey string
	Expires   int64
	Secret    string
}

type SessionJWTClaims struct {
	UserId             string `json:"user_id"`
	OpenId             string `json:"open_id"`
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

// JWT TODO expire out
func JWT(config JWTConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtToken := ctx.Request.Header.Get("token")
		if len(jwtToken) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 205,
				"msg":  "not jwtToken",
			})
			return
		}
		tokenString := strings.TrimSpace(jwtToken)
		sessionsClaims := &SessionJWTClaims{}
		_, err := jwt.ParseWithClaims(tokenString, sessionsClaims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(config.Secret), nil
			} else {
				return nil, fmt.Errorf("parse JTW error")
			}
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 206,
				"msg":  fmt.Sprintf("Parse Claims Error:%s", err),
			})
			return
		}
		ctx.Set(config.ClaimsKey, sessionsClaims)
	}
}
