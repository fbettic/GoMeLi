package main

import (
	"github.com/Chino976/GoMeLi/api_back"
	"github.com/gin-gonic/gin"
)

func Router(){

	r := gin.Default()

	//***************** BACK *****************************
	r.GET(	"/gomeli/oauth", api_back.GetCode)
	r.GET(	"/gomeli/home", api_back.Home)
	r.GET(	"/gomeli/export",api_back.Export)
	r.POST(	"/gomeli/additem", api_back.AddItem)
	r.POST(	"/gomeli/answer", api_back.Answer)


	r.Run(":8080")
}
