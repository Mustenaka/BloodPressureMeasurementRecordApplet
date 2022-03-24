package jwt

import (
	"BloodPressure/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// CustomClaims 在标准声明中加入用户id
type CustomClaims struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}

// 创建申明
func BuildClaims(exp time.Time, uid uint) *CustomClaims {
	return &CustomClaims{
		UserId: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.GlobalConfig.BasicinfoConfig.AppName,
		},
	}
}

// GenToken 生成jwt token
func GenToken(c *CustomClaims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, err := token.SignedString([]byte(secretKey))
	return ss, err
}

// ParseToken 解析jwt token
func ParseToken(jwtStr, secretKey string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(jwtStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, err
	} else {
		return nil, err
	}
}
