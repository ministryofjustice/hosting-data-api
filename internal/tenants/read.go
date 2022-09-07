package tenants

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) ReadTenant(c *gin.Context) {
	id := c.Param("id")

	for _, tenant := range tenants {
		if tenant.ID == id {
			c.JSON(http.StatusOK, tenant)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Not Found",
	})
}
