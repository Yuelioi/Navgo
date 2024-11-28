package middleware

import (
	"backend/internal/common/constants"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// JWT 中间件
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 从请求头中获取 token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			httpx.Error(w, fmt.Errorf("authorization header is missing"))
			return
		}

		// 提取 token 字符串
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			httpx.Error(w, fmt.Errorf("token is missing"))
			return
		}

		// 解析 JWT
		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return constants.ConfInst.Auth.AccessSecret, nil
		})

		if err != nil {
			httpx.Error(w, fmt.Errorf("invalid token: %v", err))
			return
		}

		// 验证 token 是否有效
		if !parsedToken.Valid {
			httpx.Error(w, fmt.Errorf("invalid token"))
			return
		}

		// 提取 Claims
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			httpx.Error(w, fmt.Errorf("invalid claims"))
			return
		}

		// 将用户名放入上下文（可根据需求放更多信息）
		username, ok := claims["username"].(string)
		if !ok {
			httpx.Error(w, fmt.Errorf("username not found in token"))
			return
		}

		// 将用户名信息存入请求上下文
		ctx := r.Context()
		ctx = context.WithValue(ctx, "username", username)
		r = r.WithContext(ctx)

		// 调用下一个处理器
		next.ServeHTTP(w, r)
	})
}
