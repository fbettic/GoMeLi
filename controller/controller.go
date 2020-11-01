package controller

import "github.com/gin-gonic/gin"

func Controller(){
	r := gin.Default()
	r.GET("/webtest/oauth.html", getCode)
	r.Run("localhost:80")
}
