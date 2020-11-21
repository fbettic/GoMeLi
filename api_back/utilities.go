package api_back

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//***********************************************************************************************
//											STRUCTS
//***********************************************************************************************

// LISTADO DE ITEMS --------------------------------------------
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


// LISTADO PREGUNTAS -------------------------------------------
type QuestsResp struct {
	Questions []question 	`json:"questions"`
	Total int				`json:"total"`
}

type question struct {
	DataCreated string		`json:"data_created"`
	ItemId string			`json:"item_id"`
	Text string				`json:"text"`
	Status string			`json:"status"`
}

// LISTADO DE PRODUCTOS VENDIDOS -------------------------------
type soldResp struct {
	Results []result		`json:"results"`
}

type result struct {
	Payments []payment		`json:"payments"`
	OrderItems []OrderItem	`json:"order_items"`
}

type payment struct {
	Reason string			`json:"reason"`
	TotalPaidAmount float64	`json:"total_paid_amount"`
	DateApproved string		`json:"date_approved"`
	Id int					`json:"id"`

}

type OrderItem struct {
	Item Item				`json:"item"`
	Quantity int			`json:"quantity"`
	UnitPrice float64		`json:"unit_price"`
}

type Item struct {
	Id string				`json:"id"`
}

type soldItem struct {
	Item string
	ItemId string
	UnitPrice float64
	Quantity int
	TotalPaidAmount float64
	TransactionID int
	TransactionDate string
}

//***********************************************************************************************
//											FUNCIONES
//***********************************************************************************************

// Esta funcion se utiliza para identificar si el token a expirado
// se utiliza comparando el body de cada request, si son iguales entonces el token ha
// expirado por lo que es necesario realizar un refresh token
func tokenError () string {
	return "{\"message\":\"Invalid token\",\"cause\":[],\"error\":\"not_found\",\"status\":401}"
}

// Funcion para abrir los archivos json
func openJson(filename string) string {

	// lee el la plantilla .json y lo almacena en un string
	jsonData, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return ""
	}

	return  string(jsonData)
}

func showResp( data []byte) string {
	var viewReq bytes.Buffer

	err := json.Indent(&viewReq, data, "", "\t")

	if err != nil {
		fmt.Errorf("Error: ",err.Error())
		return ""
	}
	return string(viewReq.Bytes())
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
		fmt.Println("Error hay que actualizar token, porfavor intente nuevamente")
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

func getSoldResp () (soldResp,error){

	resp, err := http.Get("https://api.mercadolibre.com/orders/search?seller=" +
		strconv.Itoa(User.Id) +
		"&order.status=paid&access_token=" +
		User.AccessToken)

	if err != nil {
		return soldResp{}, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return soldResp{},err
	}

	if string(data) == tokenError() {
		fmt.Println("Error hay que actualizar token, porfavor intente nuevamente")
		tokenRequest(false)
		return soldResp{}, err
	}

	var resultAux soldResp

	err = json.Unmarshal(data, &resultAux)

	if err != nil {
		return soldResp{}, err
	}

	return resultAux, nil
}

func toSoldItems(resp soldResp) []soldItem {
	var soldItemList []soldItem
	for i := 0; i< len(resp.Results); i++ {

		aux := soldItem{resp.Results[i].Payments[0].Reason,
			resp.Results[i].OrderItems[0].Item.Id,
			resp.Results[i].OrderItems[0].UnitPrice,
			resp.Results[i].OrderItems[0].Quantity,
			resp.Results[i].Payments[0].TotalPaidAmount,
			resp.Results[i].Payments[0].Id,
			resp.Results[i].Payments[0].DateApproved,
		}

		soldItemList = append(soldItemList, aux)
	}

	return soldItemList
}