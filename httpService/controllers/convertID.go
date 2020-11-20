package controllers

import (
	"net/http"
	"regexp"

	"github.com/zknow/parkingChargeAdapter/httpService/errorCode"
	"github.com/zknow/parkingChargeAdapter/httpService/service"

	"github.com/gin-gonic/gin"
)

//ConvertID 提供xps轉換卡號
func ConvertID(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.ParseParameterError})
		return
	}

	f := c.Request.PostForm

	if _, ok := f["id"]; !ok {
		c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.ParameterError})
		return
	}

	cardNumber := f.Get("id")

	regex := regexp.MustCompile(`[\f\t\n\r\v\123\x7F\x{10FFFF}\\\^\$\.\*\+\?\{\}\(\)\[\]\|G-Zg-z]`)
	if regex.MatchString(cardNumber) {
		c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.NoneCardNumberForHex})
		return
	}

	hexNum := service.ReversString(cardNumber)
	dec, err := service.HexToDecimal(hexNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"retCode": "0", "retMsg": errorCode.CardNumberParseError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"retCode": "1", "retVal": dec})
}
