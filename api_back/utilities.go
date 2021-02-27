package api_back

import (
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

// ESTRUCTURA PARA ENVIAR AL HOME
type HomeStruct struct {
	SoldItemList []soldItem 		`json:"sold_item_list"`
	ItemList []itemData				`json:"item_list"`
	QuestList []UnansweredQuest		`json:"quest_list"`
}

// LISTADO DE ITEMS --------------------------------------------
type itemsIdList struct {
	Results []string 				`json:"results"`
}

// esta estructura guarda la respuesta de meli
type itemDataReq struct{
	Title string					`json:"title"`
	AvailableQuantity int			`json:"available_quantity"`
	Price float64					`json:"price"`
	Pictures []struct{
		Url string					`json:"url"`
	}								`json:"pictures"`
}

// Esta estructura guarda los datos de los items
type itemData struct {
	Title string
	Quantity int
	Price float64
	Picture string
}


// LISTADO PREGUNTAS -------------------------------------------

// esta estructura guarda la respuesta de meli
type QuestsReq struct {
	Questions []struct{
		DataCreated string			`json:"data_created"`
		ItemId string				`json:"item_id"`
		Text string					`json:"text"`
		Status string				`json:"status"`
	} 								`json:"questions"`
	Total int						`json:"total"`
}

// Esta estructura guarda los datos de los items
type UnansweredQuest struct{
	ItemId string					`json:"item_id"`
	Text string						`json:"text"`
}


// LISTADO DE PRODUCTOS VENDIDOS -------------------------------

// esta estructura guarda la respuesta de meli
type soldReq struct {
	Results []struct{
		Payments []struct{
			Reason string			`json:"reason"`
			TotalPaidAmount float64	`json:"total_paid_amount"`
			DateApproved string		`json:"date_approved"`
			Id int					`json:"id"`
		}							`json:"payments"`
		OrderItems []struct{
			Item struct{
				Id string			`json:"id"`
			}						`json:"item"`
			Quantity int			`json:"quantity"`
			UnitPrice float64		`json:"unit_price"`
		}							`json:"order_items"`
	}								`json:"results"`
}

// Esta estructura guarda los datos de los items
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

//-------------------------Funciones De ItemList -------------------------
// esta funcion obtiene todos los ID de los items del vendedor
func getItemsID() ([]string, error) {

	// ***************** Request de items *********************
	// 1. creamos la request
	req, _ := http.NewRequest(http.MethodGet, "https://api.mercadolibre.com/users/" +
		strconv.Itoa(User.Id) +
		"/items/search", nil)

	// 2. Añadimos el access token al header
	req.Header.Add("Authorization","Bearer " + User.AccessToken)

	// 3. creamos el cliente
	client := &http.Client{}

	// 4. realizamos al consulta
	resp, err := client.Do(req)
	//*********************************************************

	if err != nil {
		// si se presento un error devolvemos un string vacio y el error obtenido
		return make([]string,0), err
	}

	// cerramos el body de la respuesta
	defer resp.Body.Close()

	// leemos la respuesta de MeLi
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// si se presento un error devolvemos un string vacio y el error obtenido
		return make([]string,0), err
	}

	// Se guardan los id de los productos
	var result itemsIdList
	err = json.Unmarshal(data, &result)


	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		return make([]string,0), err
	}

	// se devuelve la lista de IDs
	return result.Results, nil
}

// Se piden los datos de cada uno de los ID y se crea la lista de productos
func getItemsList( idList []string) ([]itemData, error){

	var itemList []itemData

	for i:=0; i < len(idList); i++ {

		item,err := getItemData(idList[i])

		if err != nil {
			return []itemData{}, errors.New("error loading item data")
		}
		// 1. Pedimos los datos del item pasando su id a getItemDataReq() y ademas
		//    Transformamos la estructura obtenida de MeLi a nuestro Propio struct
		// 2. Agregamos la estructura ya convertida en array itemData
		itemList = append(itemList, item)
	}

	// Devolvemos la lista de productos
	return itemList, nil
}

// se obtienen los datos del id que especifiquemos
func getItemData( itemID string ) (itemData,error) {

	// Pedimos los datos del item
	resp, err := http.Get("https://api.mercadolibre.com/items/" + itemID)

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		return itemData{}, err
	}

	defer resp.Body.Close()

	// leemos la respuesta de MeLi
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		return itemData{}, err
	}

	var result itemDataReq // creamos la estructura donde se guardara la respuesta de meli
	err = json.Unmarshal(data, &result) //guardamos los datos de MeLi en un struct

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		return itemData{},err
	}

	// convertimos datos de la respuesta de MeLi
	// a nuestro propio struct de datos para facilitar su manejo
	item := itemData{ result.Title,
		result.AvailableQuantity,
		result.Price,
		result.Pictures[0].Url}

	// se devuelve el struct con los datos del item
	return item, nil
}
//-------------------------------------------------------------------

//-------------------------- funciones de SoldList --------------------------
func getSoldItems () ([]soldItem,error){

	resp, err := http.Get("https://api.mercadolibre.com/orders/search?seller=" +
		strconv.Itoa(User.Id) +
		"&order.status=paid&access_token=" +
		User.AccessToken)

	if err != nil {
		return []soldItem{}, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		return []soldItem{},err
	}

	var respAux soldReq

	err = json.Unmarshal(data, &respAux)

	if err != nil {
		return []soldItem{}, err
	}

	var soldItemList []soldItem

	for i := 0; i< len(respAux.Results); i++ {

		aux := soldItem{respAux.Results[i].Payments[0].Reason,
			respAux.Results[i].OrderItems[0].Item.Id,
			respAux.Results[i].OrderItems[0].UnitPrice,
			respAux.Results[i].OrderItems[0].Quantity,
			respAux.Results[i].Payments[0].TotalPaidAmount,
			respAux.Results[i].Payments[0].Id,
			respAux.Results[i].Payments[0].DateApproved,
		}

		soldItemList = append(soldItemList, aux)
	}

	return soldItemList,nil
}
//---------------------------------------------------------------------