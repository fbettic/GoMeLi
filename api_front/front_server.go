package api_front

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidationPage( c * gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func HomePage( c * gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func NewProductPage( c * gin.Context) {
	c.HTML(http.StatusOK, "newproduct.html", nil)
}