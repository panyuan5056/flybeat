package core

import (
	"github.com/gin-gonic/gin"
)

func InitApiv1(apiv1 *gin.RouterGroup) *gin.RouterGroup {
	{

		apiv1.POST("/add", Add)
		apiv1.POST("/grok/info", GrokInfo)
		apiv1.POST("/test/data", TestData)
		apiv1.POST("/test/filter", TestFilter)
		apiv1.POST("/test/match", TestMatch)
	}
	return apiv1
}
