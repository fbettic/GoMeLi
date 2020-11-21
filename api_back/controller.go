package api_back

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
	"io/ioutil"
	"net/http"
	"strconv"
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
	value, err = sjson.Set(value, "price", c.PostForm("price"))
	value, err = sjson.Set(value, "available_quantity", c.PostForm("available_quantity"))
	value, err = sjson.Set(value, "condition", c.PostForm("condition"))

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

	if string(data) == tokenError() {
		fmt.Println("Error hay que actualizar token, porfavor intente nuevamente")
		tokenRequest(false)
		return
	}

	viewReq := showResp(data)

	// le informamos al cliente que el producto ha sido publicado con exito
	c.String(http.StatusOK, "Successfully published product\n \"meli_response\":\n %+v", viewReq)
}

// Obtener la lista de productos
func ItemList(c*gin.Context)  {
	idList, err := getItemsID()

	if err != nil{
		if err == errors.New("token has expired"){
			c.String(401, "An error has occurred, please try again")
		}
		fmt.Errorf("Error ",err.Error())
	}

	itemList := getItemsList(idList)

	data,_ := json.Marshal(itemList)

	// Solo para motrar bien los datos
	viewReq := showResp(data)
	// le informamos al cliente que el producto ha sido publicado con exito
	c.String(http.StatusOK, "{\n\"item_list\":\n %+v}", viewReq)
}

// obtener listado de preguntas
func QuestList (c*gin.Context){

	resp, err := http.Get("https://api.mercadolibre.com/questions/search?seller_id=" +
		strconv.Itoa(User.Id) + "&access_token=" + User.AccessToken +
		"&status=UNANSWERED&sort_fields=item_id,date_created&sort_type=ASC")

	if err != nil {
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	fmt.Println(string(data))
	if string(data) == tokenError() {
		fmt.Println("Error hay que actualizar token, porfavor intente nuevamente")
		tokenRequest(false)
		return
	}

	var respAux QuestsResp

	err = json.Unmarshal(data, &respAux)

	if err != nil {
		return
	}

	data,_ = json.Marshal(respAux)

	// Solo para motrar bien los datos
	viewReq := showResp(data)

	// le informamos al cliente que el producto ha sido publicado con exito
	c.String(http.StatusOK, "\"sold_list\": \n %+v ", viewReq)
	return
}

// obtener listado de productos vendidos
func SoldList(c*gin.Context){

	soldRespList, err := getSoldResp()
	soldItemList := toSoldItems(soldRespList)

	if err != nil {
		fmt.Errorf("Error: %+v ",err.Error())
		c.String(http.StatusInternalServerError, "Error: %+v", err)
	}

	data,_ := json.Marshal(soldItemList)

	// Solo para motrar bien los datos
	viewReq := showResp(data)

	// le informamos al cliente que el producto ha sido publicado con exito
	c.String(http.StatusOK, "{\n\"sold_list\":\n %+v }", viewReq)
}
