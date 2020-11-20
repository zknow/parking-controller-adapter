package httpService

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/zknow/parkingChargeAdapter/httpService/routes"
	"github.com/zknow/parkingChargeAdapter/httpService/service"

	"github.com/gin-gonic/gin"
)

//InitLog 初始化Gin log
func InitLog(f *os.File) io.Writer {
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.SetOutput(gin.DefaultWriter)
	return gin.DefaultWriter
}

//Serve 對外的api服務
func Serve(eventPool chan map[string]interface{}) {
	service.InitEventManager(eventPool)

	r := routes.CreateRouter()
	log.Fatal(r.Run(":3001"))
}
