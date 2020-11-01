package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var code string

//https://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=6719038448258240&redirect_uri=https://c8b717f929bc.ngrok.io/oauth/code/


type Token struct {
	Grant_type string 		`json:"grant_type"`
	Client_id int 			`json:"client_id"`
	Client_secret string 	`json:"client_secret"`
	Code string				`json:"code"`
	Redirect_uri string 	`json:"redirect_uri"`
}

type TokenResp struct {
	Access_token string 	`json:"access_token"`
	Token_type string 		`json:"token_type"`
	Expires_in int 			`json:"expires_in"`
	Scope string 			`json:"scope"`
	User_id int 			`json:"user_id"`
	Refresh_token string 	`json:"refresh_token"`
}


func getCode(c *gin.Context){
	code = c.Query("code")
	tokenRequest()
}

func tokenRequest() {

	u := Token{	Grant_type: "authorization_code",
				Client_id: 6719038448258240,
				Client_secret: "qmxiwj6zMUkNyWs1YzdOHkuCkkquJfVw",
				Code: code,
				Redirect_uri: "https://568130c7da0a.ngrok.io/webtest/oauth.html"}

	b, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	resp, err := http.Post("https://api.mercadolibre.com/oauth/token","application/json; application/x-www-form-urlencoded", bytes.NewBuffer(b))

	fmt.Println("Error",err)

	if err != nil {
		fmt.Errorf("Error",err.Error())
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	bodyString := string(data)
	fmt.Println(bodyString)

	var tokenResp TokenResp
	json.Unmarshal(data, &tokenResp)
	fmt.Printf("%+v\n", tokenResp)
}



