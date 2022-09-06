package tenants

import (
	"github.com/gin-gonic/gin"
)

type handler struct{}

func RegisterRoutes(r *gin.Engine) {
	h := &handler{}

	routes := r.Group("/tenants")
	routes.GET("/", h.ListTenants)
	routes.GET("/:id", h.ReadTenant)
}
