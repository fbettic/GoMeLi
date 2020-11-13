package api_back

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Esta estructura guarda la respuesta de meli
type resultsList struct {
	Results []string 		`json:"results"`
}

type itemDataReq struct{
	Title string			`json:"title"`
	AvailableQuantity int	`json:"available_quantity"`
	Price float64			`json:"price"`
	Pictures []pictures		`json:"pictures"`
}
type pictures struct {
	Url string				`json:"url"`
}

// Esta estructura guarda los datos de los items
type itemData struct {
	Title string
	Quantity int
	Price float64
	Picture string
}

func ItemList(c*gin.Context)  {
	idList, err := getItemsID()

	if err != nil{
		if err.Error() == "token has expired" {
			c.String(401, "Token vencido. Por favor recargar la página")
		}
		fmt.Errorf("Error ",err.Error())
	} else {
		itemList := getItemsList(idList)

		data,_ := json.Marshal(itemList)

		// Solo para motrar bien los datos
		viewReq := showResp(data)
		// le informamos al cliente que el producto ha sido publicado con exito
		c.String(http.StatusOK, "{\n\"item_list\":\n %+v}", viewReq)
	}
}

// esta funcion obtiene todos los ID de los items del vendedor
func getItemsID() ([]string, error) {

	// se envia la peticion get para obtener los datos del vendedor
	resp, err := http.Get("https://api.mercadolibre.com/users/" +
		strconv.Itoa(User.Id) +
		"/items/search?&access_token=" +
		User.AccessToken)

	if err != nil {
		return make([]string,0), err
	}

	// cerramos el body de la respuesta
	defer resp.Body.Close()

	// leemos la respuesta de MeLi
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return make([]string,0), err
	}

	// Se verifica que el token no haya expirado
	// Si ha expirado se realiza un refresh token
	if string(data) == tokenError() {
		tokenRequest(false)
		return make([]string, 0), errors.New("token has expired")
	}

	// Se guardan los id de los productos
	var resultAux resultsList
	err = json.Unmarshal(data, &resultAux)

	if err != nil {
		fmt.Println(err)
		fmt.Errorf("Error ",err.Error())
		return make([]string,0), err
	}

	// se devuelve la lista de IDs
	return resultAux.Results, nil
}

// Se piden los datos de cada uno de los ID y se crea la lista de productos
func getItemsList( idList []string) []itemData{

	var itemList []itemData
	for i:=0; i < len(idList); i++ {

		// 1. Pedimos los datos del item pasando su id a getItemDataReq()
		// 2. Transformamos la estructura obtenida de MeLi a nuestro Propio struct con toItemData
		// 3. Agregamos la estructura ya convertida en array itemData
		itemList = append(itemList, toItemData(getItemDataReq(idList[i])))
	}

	// Devolvemos la lista de productos
	return itemList
}

// se obtienen los datos del id que especifiquemos
func getItemDataReq( itemID string ) itemDataReq{

	// Pedimos los datos del item
	resp, err := http.Get("https://api.mercadolibre.com/items/" + itemID)

	if err != nil {
		fmt.Errorf("Error ",err.Error())
		return itemDataReq{}
	}

	defer resp.Body.Close()

	// leemos la respuesta de MeLi
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Errorf("Error ",err.Error())
		return itemDataReq{}
	}

	var resultAux itemDataReq
	err = json.Unmarshal(data, &resultAux) //guardamos los datos de MeLi en un struct

	if err != nil {
		fmt.Println(err)
		fmt.Errorf("Error ",err.Error())
		return itemDataReq{}
	}

	// se devuelve el struct con los datos
	return resultAux
}

// Se convierten los struct datos de la respuesta de MeLi
// a nuestro propio struct de datos para facilitar su manejo
func toItemData(req itemDataReq) itemData{

	item := itemData{ req.Title,
					req.AvailableQuantity,
					req.Price,
					req.Pictures[0].Url}
	return item
}

