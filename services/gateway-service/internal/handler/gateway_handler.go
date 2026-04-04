package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/arch-yunus/microservices-101/proto/order"
	"github.com/arch-yunus/microservices-101/proto/product"
	"github.com/labstack/echo/v4"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayHandler struct {
	productClient product.ProductServiceClient
	orderClient   order.OrderServiceClient
	cb            *gobreaker.CircuitBreaker
}

func NewGatewayHandler() *GatewayHandler {
	productSvcAddr := os.Getenv("PRODUCT_SERVICE_ADDR")
	if productSvcAddr == "" {
		productSvcAddr = "product-service:50051"
	}

	orderSvcAddr := os.Getenv("ORDER_SERVICE_ADDR")
	if orderSvcAddr == "" {
		orderSvcAddr = "order-service:50052"
	}

	// Internal gRPC Connections
	pConn, _ := grpc.Dial(productSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	oConn, _ := grpc.Dial(orderSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// Circuit Breaker Settings
	cbSettings := gobreaker.Settings{
		Name: "GatewayBreaker",
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 3
		},
	}

	return &GatewayHandler{
		productClient: product.NewProductServiceClient(pConn),
		orderClient:   order.NewOrderServiceClient(oConn),
		cb:            gobreaker.NewCircuitBreaker(cbSettings),
	}
}

func (h *GatewayHandler) GetProduct(c echo.Context) error {
	id := c.Param("id")
	
	resp, err := h.cb.Execute(func() (interface{}, error) {
		return h.productClient.GetProduct(context.Background(), &product.GetProductRequest{Id: id})
	})

	if err != nil {
		if err == gobreaker.ErrOpenState {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Ürün servisi şu anda meşgul (Circuit Breaker OPEN)"})
		}
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Ürün bulunamadı"})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayHandler) CreateOrder(c echo.Context) error {
	var req struct {
		ProductID string `json:"product_id"`
		Quantity  int32  `json:"quantity"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Geçersiz istek"})
	}

	resp, err := h.cb.Execute(func() (interface{}, error) {
		return h.orderClient.CreateOrder(context.Background(), &order.CreateOrderRequest{
			ProductId: req.ProductID,
			Quantity:  req.Quantity,
		})
	})

	if err != nil {
		if err == gobreaker.ErrOpenState {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Sipariş servisi şu anda meşgul (Circuit Breaker OPEN)"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Sipariş oluşturulamadı"})
	}

	return c.JSON(http.StatusCreated, resp)
}
