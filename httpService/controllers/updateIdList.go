package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zknow/parkingChargeAdapter/httpService/errorCode"
	"github.com/zknow/parkingChargeAdapter/httpService/service"
)

//UpdateIdList 更新的參數
func UpdateIdList(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.ParseParameterError})
		return
	}

	reqForm := c.Request.PostForm
	if _, ok := reqForm["update"]; !ok {
		c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.ParameterError})
		return
	}

	updateIdList := reqForm.Get("update")

	c.JSON(http.StatusOK, gin.H{"retCode": "1", "retMsg": "ok", "retVal": updateIdList})

	service.EventMgr.PushEvent("XPS", "UpdateIdList", updateIdList)
}
