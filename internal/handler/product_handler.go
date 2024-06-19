package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/omarelweshy/EcomMaster-product-service/internal/service"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	return
}
