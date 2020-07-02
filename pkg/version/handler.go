package version

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVersion(c *gin.Context) {
	info := Get()
	infoJson, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err)
	}

	c.String(http.StatusOK, string(infoJson))
}
