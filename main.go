package main

import (
	"GoMeLI/controller"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	controller.Controller()
}
