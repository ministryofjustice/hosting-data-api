package tenants

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h handler) ReadTenant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		for _, tenant := range tenants {
			if tenant.ID == id {
				c.JSON(http.StatusOK, tenant)
				return
			}
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Not Found",
	})
}
