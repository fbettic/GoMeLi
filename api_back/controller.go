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

	// Abrimos el archivo .json que contiene todos los datos que MeLi requiere para publicar
	// (una especie de platilla, para evitar los structs)
	jsonData := openJson("./api_back/json/addItem.json")

	//esto es para la peticion que vendra del front, a implementar
	//**************************************************************
	//data,err := ioutil.ReadAll(c.Request.Body)
	//
	//if err != nil {
	//	fmt.Println("Error: ", err.Error())
	//	return
	//}
	//
	//fmt.Println(string(data))

	//c.JSON(http.StatusOK,c)
	//*************************************************************

	// Cambiamos los valores de la plantilla .json por los que envio el ususario
	// en este caso cambiamos la variable precio (libreria: tidwall/sjson)
	value, _ := sjson.Set(jsonData, "price", "8000")

	println(value)

	// combertimos la variable value (string que el json a enviar ya modificado) a un array de bytes para poder
	// realizar el metodo post
	b := []byte(value)


	fmt.Println(string(b))

	// realizamos el post
	resp, err := http.Post("https://api.mercadolibre.com/items?access_token=" + AccessToken,
		"application/json; application/x-www-form-urlencoded",
		bytes.NewBuffer(b))

	if err != nil {
		fmt.Errorf("Error: ",err.Error())
		return
	}

	defer resp.Body.Close()

	// leemos la respuesta de MeLi
	data, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))

}

func openJson(filename string) string {

	// lee el la plantilla .json y lo almacena en un string
	jsonData, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return ""
	}

	return  string(jsonData)
}