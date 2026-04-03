package auth

import (
	"errors"
	"net/http"
	"strings"
)

// JWTAuthMiddleware basit bir token kontrol mekanizmas (Eitim Amacl)
// Gerçek projelerde "github.com/golang-jwt/jwt" kütüphanesi kullanlmaldr.
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Yetkisiz: Token bulunamadı", http.StatusUnauthorized)
			return
		}

		// "Bearer secret-token-123" formatn kontrol et
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" || parts[1] != "secret-token-123" {
			http.Error(w, "Yetkisiz: Gecersiz token", http.StatusUnauthorized)
			return
		}

		// Token geçerliyse bir sonraki handler'a geç
		next.ServeHTTP(w, r)
	})
}
