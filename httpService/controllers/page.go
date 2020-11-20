package controllers

import (
	"net/http"

	"github.com/zknow/parkingChargeAdapter/httpService/errorCode"
	"github.com/zknow/parkingChargeAdapter/httpService/service"

	"github.com/gin-gonic/gin"
)

//Page Screen 螢幕通知車柱
func Page(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.ParseParameterError})
		return
	}

	f := c.Request.PostForm
	switch f.Get("action") {
	case "carNumber":
		service.EventMgr.PushEvent("Screen", "CarNumber", map[string]string{"Number": f.Get("number")})
	case "carPlat":
		service.EventMgr.PushEvent("Screen", "CarPlat", map[string]string{"Number": f.Get("number"), "AccessTime": f.Get("accessTime")})
	case "invoice":
		service.EventMgr.PushEvent("Screen", "Invoice", map[string]string{"SaveInvoice": f.Get("saveInvoice"), "TxNumber": f.Get("txNumber")})
	case "emergency":
		service.EventMgr.PushEvent("Screen", "Emergency", "emergency")
	default:
		c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.ParameterError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"retCode": "1", "retMsg": "ok"})
}
