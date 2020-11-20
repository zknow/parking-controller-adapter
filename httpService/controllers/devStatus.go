package controllers

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/zknow/parkingChargeAdapter/httpService/errorCode"
	"github.com/zknow/parkingChargeAdapter/httpService/service"

	"github.com/gin-gonic/gin"
)

//DevStatus 查詢修改車柱狀態
func DevStatus(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.ParseParameterError})
		return
	}

	reqForm := c.Request.PostForm

	filename := "/tmp/devStatus"

	switch reqForm.Get("action") {
	case "alive":
		service.EventMgr.PushEvent("XPS", "DevStatus", "alive")
		err := ioutil.WriteFile(filename, []byte("alive"), os.ModePerm)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.SetStatusError})
		}
	case "stop":
		service.EventMgr.PushEvent("XPS", "DevStatus", "stop")
		err := ioutil.WriteFile(filename, []byte("stop"), os.ModePerm)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.SetStatusError})
		}
	case "getStatus":
		if _, statErr := os.Stat(filename); !os.IsNotExist(statErr) {
			bs, _ := ioutil.ReadFile(filename)
			c.JSON(http.StatusOK, gin.H{"retCode": "1", "status": string(bs)})
		} else {
			c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.LostStatusFile})
		}
		return
	default:
		c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.ParameterError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"retCode": "1", "retMsg": "ok"})
}
