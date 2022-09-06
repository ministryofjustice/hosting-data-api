package tenants

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type tenant struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var tenants = []tenant{
	{
		ID:    1,
		Title: "Tenant 1",
	},
	{
		ID:    2,
		Title: "Tenant 2",
	},
	{
		ID:    3,
		Title: "Tenant 3",
	},
	{
		ID:    4,
		Title: "Tenant 4",
	},
	{
		ID:    5,
		Title: "Tenant 5",
	},
}

func (h handler) ListTenants(c *gin.Context) {
	c.JSON(http.StatusOK, tenants)
}
