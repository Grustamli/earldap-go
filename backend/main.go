package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("account/get-username", requestUsername)
		v1.POST("account/reset-password", resetPassword)
		v1.POST("account/verify-otp", verifyOTP)
		v1.POST("account/set-password", setPassword)
	}
	router.Run()
}

func requestUsername(c *gin.Context) {

}

func resetPassword(c *gin.Context) {

}

func verifyOTP(c *gin.Context) {

}

func setPassword(c *gin.Context) {

}
