package api_back

import  (
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


type MessageStruct struct {
	Message string					`json:"message"`
	Status int 						`json:"status"`
}

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
	Id string						`json:"id"`
	Title string					`json:"title"`
	InitialQuantity int				`json:"initial_quantity"`
	SoldQuantity int				`json:"sold_quantity"`
	Price float64					`json:"price"`
	Pictures []struct{
		Url string					`json:"url"`
	}								`json:"pictures"`
}

// Esta estructura guarda los datos de los items
type itemData struct {
	Id string						`json:"id"`
	Title string					`json:"title"`
	Quantity int					`json:"quantity"`
	Price float64					`json:"price"`
	Picture string					`json:"picture"`
}


// LISTADO PREGUNTAS -------------------------------------------

// esta estructura guarda la respuesta de meli
type QuestsReq struct {
	Questions []struct{
		DataCreated string			`json:"data_created"`
		ItemId string				`json:"item_id"`
		Text string					`json:"text"`
		Status string				`json:"status"`
		Id int 						`json:"id"`
	}								`json:"questions"`
	Total int						`json:"total"`
}

// Esta estructura guarda los datos de los items
type UnansweredQuest struct{
	ItemTitle string				`json:"item_title"`
	ItemId string					`json:"item_id"`
	Id int							`json:"id"`
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
	Item string						`json:"item"`
	ItemId string					`json:"item_id"`
	UnitPrice float64				`json:"unit_price"`
	Quantity int					`json:"quantity"`
	TotalPaidAmount float64			`json:"total_paid_amount"`
	TransactionID int				`json:"transaction_id"`
	TransactionDate string			`json:"transaction_date"`
}

//***********************************************************************************************
//											FUNCIONES
//***********************************************************************************************



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
func getItemsID(user ReqUserData) ([]string, error) {

	// ***************** Request de items *********************
	// 1. creamos la request
	req, _ := http.NewRequest(http.MethodGet, "https://api.mercadolibre.com/users/" +
		strconv.Itoa(user.IdMeli) +
		"/items/search", nil)

	// 2. Añadimos el access token al header
	req.Header.Add("Authorization","Bearer " + user.AccessToken)

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

	var err1 error
	err1 = nil

	chit := make(chan itemData)
	cherr := make(chan error)

	for i:=0; i < len(idList); i++ {

		go func(chit chan itemData, cherr chan error,stid string){

			item,err := getItemData(stid)

			if err != nil {
				chit <- item
				cherr <- errors.New("error loading item data")
				return
			}
			// 1. Pedimos los datos del item pasando su id a getItemDataReq() y ademas
			//    Transformamos la estructura obtenida de MeLi a nuestro Propio struct
			// 2. Agregamos la estructura ya convertida en array itemData
			chit <- item
			cherr <- nil

		}(chit,cherr,idList[i])

	}

	for i:=0; i < len(idList); i++ {
		item := <- chit
		itemList = append(itemList,item)
		err := <- cherr
		if err != nil{
			err1 = err
		}
	}

	// Devolvemos la lista de productos
	return itemList, err1

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
	item := itemData{ 
		result.Id,
		result.Title,
		result.InitialQuantity - result.SoldQuantity,
		result.Price,
		result.Pictures[0].Url}

	// se devuelve el struct con los datos del item
	return item, nil
}
//-------------------------------------------------------------------

//-------------------------- funciones de SoldList --------------------------
func getSoldItems (user ReqUserData) ([]soldItem,error){

	resp, err := http.Get("https://api.mercadolibre.com/orders/search?seller=" +
		strconv.Itoa(user.IdMeli) +
		"&order.status=paid&access_token=" +
		user.AccessToken)

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

//-------------------- testea el tocken para saber si esta vencido ----------------------------
func testToken(id string) (ReqUserData,error){

	user,err := obtenerDatosUsuario(id)

	if err != nil {
		fmt.Println(err)
		return ReqUserData{},errors.New("error al cargar los datos de usuario 1")
	}

	tokenError := "{\"message\":\"invalid_token\",\"error\":\"not_found\",\"status\":401,\"cause\":[]}"

	req, _ := http.NewRequest(http.MethodGet, "https://api.mercadolibre.com/users/" +
		strconv.Itoa(user.IdMeli) +
		"?attributes=status", nil)

	// 2. Añadimos el access token al header
	req.Header.Add("Authorization","Bearer " + user.AccessToken)

	// 3. creamos el cliente
	client := &http.Client{}

	// 4. realizamos al consulta
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return ReqUserData{},errors.New("error al testear el token")
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))

	if string(data)==tokenError{
		tokenRequest(user.RefreshToken, false)
		user,err = obtenerDatosUsuario(id)
		if err != nil {
			fmt.Println(err)
			return ReqUserData{},errors.New("error al cargar los datos de usuario")
		}
		fmt.Println("token actualizado")
		return user,nil
	}
	return user,nil
}
//---------------------------------------------------------------------------------------------