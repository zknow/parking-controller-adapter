package routes

import (
	"github.com/zknow/parkingChargeAdapter/httpService/controllers"

	"github.com/gin-gonic/gin"
)

//CreateRouter 建立Router
func CreateRouter() *gin.Engine {
	r := gin.Default()
	// Recovery異常回覆500
	r.Use(gin.Recovery())
	r.POST("/page", controllers.Page)
	r.POST("/updateBlackWhiteList", controllers.UpdateIdList)
	r.POST("/devStatus", controllers.DevStatus)
	r.POST("/gate", controllers.Gate)
	r.POST("/counter888", controllers.Counter888)
	r.POST("/convertID", controllers.ConvertID)
	return r
}
