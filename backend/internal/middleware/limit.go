package middleware

import (
	"net/http"
	"sync"
	"time"
)

func Limit() func(next http.HandlerFunc) http.HandlerFunc {
	type accessRecord struct {
		count    int
		lastTime time.Time
	}

	accessRecords := make(map[string]accessRecord)
	var mu sync.Mutex

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ip := getIPFromRequest(r)

			mu.Lock()
			record, exists := accessRecords[ip]
			if !exists {
				record = accessRecord{count: 0, lastTime: time.Now()}
			}

			// 检查是否超过 1 分钟
			if time.Since(record.lastTime) > time.Minute {
				record.count = 0
			}

			// 检查访问次数
			if record.count >= 10 {
				http.Error(w, "Too many requests, please try again later.", http.StatusTooManyRequests)
				mu.Unlock()
				return
			}

			// 更新访问记录
			record.count++
			record.lastTime = time.Now()
			accessRecords[ip] = record
			mu.Unlock()

			// 继续处理请求
			next.ServeHTTP(w, r)
		}
	}
}

// getIPFromRequest 从请求中提取客户端的 IP 地址
func getIPFromRequest(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}
	println(ip)
	return ip
}
