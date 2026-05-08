package middleware

import (
	"net/http"
	"strings"
	"wenci/internal/consts"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(r *ghttp.Request) {
	raw := r.Header.Get("Authorization")
	tokenString := strings.TrimSpace(raw)
	if tokenString == "" {
		r.Response.WriteJson(http.StatusForbidden)
		r.Exit()
		return
	}

	// 支持两种格式：
	// 1）Authorization: Bearer <jwt>（推荐，标准格式）
	// 2）Authorization: <jwt>（兼容旧客户端）
	const bearerPrefix = "Bearer "
	if strings.HasPrefix(tokenString, bearerPrefix) {
		tokenString = strings.TrimSpace(tokenString[len(bearerPrefix):])
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})
	if err != nil || !token.Valid {
		r.Response.WriteJson(http.StatusForbidden)
		r.Exit()
		return
	}
	r.Middleware.Next()
}
