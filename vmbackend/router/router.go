package router

import (
	"vmbackend/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	vm := r.Group("/vmapi")
	{
		vm.POST("/login", controller.Login)
		vm.GET("/goods", controller.GetGoodsList)
		vm.GET("/inbound", controller.GetInboundList)
		vm.GET("/outbound", controller.GetOutboundList)
		vm.GET("/inventory", controller.GetInventoryList)
		vm.POST("/inbound", controller.PostInBound)
		vm.POST("/outbound", controller.PostOutBound)
	}
}
