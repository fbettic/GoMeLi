package api_back

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)


// URL puesta en la app de MeLi
var url string = "http://localhost:8080/gomeli/oauth"

// struct que se enviara como body para obtener el Access token
type Token struct {
	GrantType string 		`json:"grant_type"`
	ClientId int 			`json:"client_id"`
	ClientSecret string 	`json:"client_secret"`
	Code string				`json:"code,omitempty"`				// ",omitempty" sirve para que al pasar el struct a
	RefreshToken string		`json:"refresh_token,omitempty"`	// JSON, si la var esta vacia se omite y no se
	RedirectUri string 		`json:"redirect_uri,omitempty"` 	// toma en cuenta en la conversion
}

// struct para almacenar la respuesta de MeLi
type TokenResp struct {
	AccessToken string 		`json:"access_token"`
	TokenType string 		`json:"token_type"`
	ExpiresIn int 			`json:"expires_in"`
	Scope string 			`json:"scope"`
	UserId int 				`json:"user_id"`
	RefreshToken string 	`json:"refresh_token"`
}


func GetCode(c *gin.Context){

	// obtenemos el codigo de intercambio y nos aseguramos de que no este vacio
	code := c.Query("code")

	if code == "" {
		c.String(400, "HTTP 400 Missing param code")
		return
	}

	// redireccionamos a la pagina de home de nuestro front
	c.Redirect(302, "http://localhost/gomeli/home.html")

	// llamamos a la funcion para obtener el token por "primera vez"
	tokenRequest(code,true)
}

func tokenRequest( code string, firstChange bool ) {

	// pedimos los datos para crear el body que sera enviado a MeLi
	b, err := json.Marshal(bodyToken( code, firstChange ))

	// comprobamos que no haya un error en la conversion
	if err != nil {
		fmt.Println(err)
		return
	}

	// hacemos el post del body
	resp, err := http.Post("https://api.mercadolibre.com/oauth/token",
					 "application/json; application/x-www-form-urlencoded",
								bytes.NewBuffer(b))

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		return
	}

	defer resp.Body.Close()

	// leemos la respuesta de MeLi
	data, err := ioutil.ReadAll(resp.Body)

	// decodificamos la respuesta y la almacenamos en una tokenResp
	var tokenResp TokenResp

	json.Unmarshal(data, &tokenResp)

	user := ReqUserData{
		tokenResp.UserId,
		tokenResp.AccessToken,
		tokenResp.RefreshToken,
	}

	fmt.Println(string(data))

	// enviamos los nuevos datos a la base de datos
	actualizarToken(user)
}

func bodyToken( code string, firstChange bool ) Token {

	// si es la primera vez que se pide el token
	if firstChange {
		return Token{	GrantType: "authorization_code",
						ClientId: 6719038448258240,
						ClientSecret: "qmxiwj6zMUkNyWs1YzdOHkuCkkquJfVw",
						Code: code,
						RedirectUri: url}
	}

	// si es una peticion de refresh token
	return Token{	GrantType: "refresh_token",
					ClientId: 6719038448258240,
					ClientSecret: "qmxiwj6zMUkNyWs1YzdOHkuCkkquJfVw",
					RefreshToken: code}

}





