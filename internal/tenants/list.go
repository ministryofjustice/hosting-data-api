package tenants

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tenant struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	BillingStartDate string `json:"billing_start"`
	BillingEndDate   string `json:"billing_end"`
}

var tenants = []tenant{}

func (h handler) ListTenants(c *gin.Context) {
	var file, _ = ioutil.ReadFile("./data/aws/accounts.json")
	_ = json.Unmarshal([]byte(file), &tenants)

	c.JSON(http.StatusOK, tenants)
}
