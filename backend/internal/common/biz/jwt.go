package biz

import (
	"backend/internal/common/constants"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Validate(tokenString string) (string, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(constants.ConfInst.Auth.AccessSecret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}
	// 验证是否过期
	exp, ok := claims["exp"].(float64)
	if !ok {
		return "", fmt.Errorf("exp not found in token")
	}

	if time.Unix(int64(exp), 0).Before(time.Now()) {
		return "", fmt.Errorf("token has expired")
	}

	// 提取用户名或用户ID
	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username not found in token")
	}
	return username, nil
}

func Generate(username, role string) (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,                                                                                       // 用户名
		"role":     role,                                                                                           // 角色
		"exp":      time.Now().Add(time.Duration(constants.SVCInst.Config.Auth.AccessExpire) * time.Second).Unix(), // 过期时间
		"iat":      time.Now().Unix(),                                                                              // 签发时间
	})

	tokenString, err := tokens.SignedString([]byte(constants.ConfInst.Auth.AccessSecret))
	return tokenString, err
}
