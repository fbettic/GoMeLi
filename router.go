package main

import (
	"github.com/Chino976/GoMeLi/api_back"
	"github.com/Chino976/GoMeLi/api_front"
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
	r.GET("/webtest/itemlist", api_back.ItemList)
	r.GET("/webtest/soldlist", api_back.SoldList)
	r.GET("/webtest/questlist", api_back.QuestList)

	api_back.ReadUserList()
	api_back.LoadUserData(666272328)

	r.Run(":80")
}
