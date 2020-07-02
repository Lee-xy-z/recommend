package routers

import (
	v "github.com/Lee-xy-z/recommend/pkg/version"
	"github.com/gin-gonic/gin"
)

func RegisterAPI() {
	router := gin.Default()

	v2 := router.Group("/v2")
	{
		v2.GET("/version", v.GetVersion)
	}

	router.Run()
}
