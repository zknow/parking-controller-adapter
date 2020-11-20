package controllers

import (
	"net/http"

	"github.com/zknow/parkingChargeAdapter/httpService/errorCode"
	"github.com/zknow/parkingChargeAdapter/httpService/service"

	"github.com/gin-gonic/gin"
)

//Counter888 加減計數器
func Counter888(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.ParseParameterError})
		return
	}

	f := c.Request.PostForm
	service.EventMgr.PushEvent("XPS", "Counter888", map[string]string{"option": f.Get("option"), "count": f.Get("count")})
	c.JSON(http.StatusOK, gin.H{"retCode": "1", "retMsg": "ok"})
}
