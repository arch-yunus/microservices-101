package handler

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// NewGatewayHandler tum istekleri ic servislere yonlendirir
func NewGatewayHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var targetURL string

		// Basit "Path-based Routing" stratejisi
		switch {
		case strings.HasPrefix(r.URL.Path, "/products"):
			targetURL = "http://product-service:8081" // Docker daki servis ismi
		case strings.HasPrefix(r.URL.Path, "/orders"):
			targetURL = "http://order-service:8082"
		default:
			http.Error(w, "Gecersiz servis yolu", http.StatusNotFound)
			return
		}

		target, _ := url.Parse(targetURL)
		proxy := httputil.NewSingleHostReverseProxy(target)

		fmt.Printf("?? Yonlendiriliyor: %s -> %s\n", r.URL.Path, targetURL)
		proxy.ServeHTTP(w, r)
	})
}
