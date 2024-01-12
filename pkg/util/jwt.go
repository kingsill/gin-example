package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/kingsill/gin-example/pkg/setting"
)

// 加载配置文件中设置的密钥
var jwtSecret = []byte(setting.JwtSecret)

// Claims 定义claims结构体
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	//创建 CustomClaims 结构体，用来封装 jwt 信息
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	//创建 header和payload部分
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//得到完整的token字符串，这里为加入签名signature部分
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	//解码过程
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	//验证是否时间过期
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
