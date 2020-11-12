package api_back

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
	"io/ioutil"
	"net/http"
)

// Funcion para publicar items
func AddItem(c*gin.Context) {

	// Abrimos el archivo .json que contiene todos los datos que MeLi requiere para publicar un item
	// (una especie de platilla, para evitar los structs)
	jsonData := openJson("./api_back/json/addItem.json")

	//**************************************************************
	//Aqui se debe implementar el codigo para recibir los datos del front
	//**************************************************************

	fmt.Println(c.PostForm("title"))
	value, err := sjson.Set(jsonData, "title", c.PostForm("title"))

	if err != nil {
		fmt.Errorf("Error: ",err.Error())
		return
	}

	// convertimos la variable value (string con el json modificado con los datos del front)
	// a un array de bytes para luego realizar el metodo post
	b := []byte(value)

	fmt.Println(string(b))

	// realizamos el post
	resp, err := http.Post("https://api.mercadolibre.com/items?access_token=" + AccessToken,
		"application/json; application/x-www-form-urlencoded",
		bytes.NewBuffer(b))

	if err != nil {
		fmt.Errorf("Error: %v",err.Error())
		return
	}

	// cerramos el body
	defer resp.Body.Close()

	// leemos la respuesta de MeLi
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Errorf("Error: ",err.Error())
		return
	}

	viewReq := showResp(data)

	// le informamos al cliente que el producto ha sido publicado con exito
	c.String(http.StatusOK, "Successfully published product\n \"meli_response\":\n %+v", viewReq)
}