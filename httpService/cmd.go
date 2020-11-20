package httpService

import (
	"io"
	"os"

	log "github.com/gogf/gf/os/glog"

	"github.com/zknow/parkingChargeAdapter/httpService/routes"
	"github.com/zknow/parkingChargeAdapter/httpService/service"

	"github.com/gin-gonic/gin"
)

//InitLog 初始化Gin log
func InitLog(f *os.File) {
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.SetWriter(gin.DefaultWriter)
}

//Serve 對外的api服務
func Serve(eventPool chan map[string]interface{}) {
	service.InitEventManager(eventPool)

	r := routes.CreateRouter()
	log.Fatal(r.Run(":3001"))
}
