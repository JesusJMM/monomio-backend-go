package apiutils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIntQueryParam(c *gin.Context, param string, def int) (int, bool) {
	if c.Query(param) != "" {
		var err error
    p, err := strconv.Atoi(c.Query(param))
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("%s query param must be a number", param))
			return p, false
		}
	}
	return def, true
}
