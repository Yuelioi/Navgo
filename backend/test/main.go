package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("123")
var expireTime = time.Hour

func ValidateAndGetUser(tokenString string) (bool, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecret, nil
	})
	if err != nil {
		return false, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return false, fmt.Errorf("invalid claims")
	}
	// 验证是否过期
	exp, ok := claims["exp"].(float64)
	if !ok {
		return false, fmt.Errorf("exp not found in token")
	}

	if time.Unix(int64(exp), 0).Before(time.Now()) {
		return false, fmt.Errorf("token has expired")
	}

	// 提取用户名或用户ID
	_, ok = claims["username"].(string)
	if !ok {
		return false, fmt.Errorf("username not found in token")
	}
	return true, nil
}

func Generate(username, role string) (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,                          // 用户名
		"role":     role,                              // 角色
		"exp":      time.Now().Add(expireTime).Unix(), // 过期时间
		"iat":      time.Now().Unix(),                 // 签发时间
	})

	tokenString, err := tokens.SignedString(jwtSecret)
	return tokenString, err
}

func main() {

	token, err := Generate("zhangsan", "admin")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)
	ok, _ := ValidateAndGetUser(token)
	print(ok)
}
