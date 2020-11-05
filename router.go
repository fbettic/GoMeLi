package main

import (
	"GoMeLi/api_back"
	"GoMeLi/api_front"
	"github.com/gin-gonic/gin"
)

func Router(){

	r := gin.Default()

	r.LoadHTMLGlob("public/*.html")

	//***************** FRONT *****************************
	r.GET("/webtest", api_front.ValidationPage)
	r.GET("/webtest/home.html", api_front.HomePage)
	r.GET("/webtest/newproduct.html", api_front.NewProductPage)

	//***************** BACK *****************************
	r.GET("/webtest/oauth", api_back.GetCode)
	r.POST("/webtest/additem", api_back.AddItem)


	r.Run(":80")
}
