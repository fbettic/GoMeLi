package api_back_test

import (
	"github.com/Chino976/GoMeLi/api_back"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	//***************** BACK *****************************
	r.GET(	"/gomeli/oauth", api_back.GetCode)
	r.GET(	"/gomeli/home", api_back.Home)
	r.GET(	"/gomeli/export",api_back.Export)
	r.POST(	"/gomeli/additem", api_back.AddItem)
	r.POST(	"/gomeli/answer", api_back.Answer)
	return r
}

func TestExport(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/gomeli/export?id=1", nil)
	router.ServeHTTP(w, req)

	expected, _ := ioutil.ReadFile("./json/response.json")

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}