package controllers

import (
	"net/http"

	"github.com/zknow/parkingChargeAdapter/httpService/errorCode"
	"github.com/zknow/parkingChargeAdapter/httpService/service"

	"github.com/gin-gonic/gin"
)

//Gate 遠端開門功能
func Gate(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.ParseParameterError})
		return
	}

	f := c.Request.PostForm
	if act := f.Get("action"); act != "open" && act != "keepOpen" && act != "close" {
		c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.ParameterError})
		return
	}
	service.EventMgr.PushEvent("XPS", "Gate", map[string]string{"action": f.Get("action")})
	c.JSON(http.StatusOK, gin.H{"retCode": "1", "retMsg": "ok"})
}
