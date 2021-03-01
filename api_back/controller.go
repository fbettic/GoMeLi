package api_back

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
	"io/ioutil"
	"net/http"
	"strconv"
)



// Funcion para publicar items
func AddItem(c*gin.Context) {

	id := c.Query("id")

	user,err := testToken(id)

	if err != nil {
		fmt.Println(err)
		response := MessageStruct{"Error al tratar de obtener los datos de usuario", 401}
		c.JSON(401,response)
		return
	}

	// Abrimos el archivo .json que contiene todos los datos que MeLi requiere para publicar un item
	// (una especie de platilla, para evitar el exeso de structs)
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
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		return
	}

	// convertimos la variable value (string con el json modificado con los datos del front)
	// a un array de bytes para luego realizar el metodo post
	b := []byte(value)

	fmt.Println(string(b))



	// 1. creamos la request
	req, _ := http.NewRequest(http.MethodPost, "https://api.mercadolibre.com/items", bytes.NewBuffer(b))

	// 2. Añadimos el access token al header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization","Bearer " + user.AccessToken)

	// 3. creamos el cliente
	client := &http.Client{}

	// 4. realizamos al consulta
	resp, err := client.Do(req)

	/* realizamos el post de los datos a mercado libre
	resp, err := http.Post("https://api.mercadolibre.com/items?access_token=" + user.AccessToken,
		"application/json; application/x-www-form-urlencoded",
		bytes.NewBuffer(b))
	*/


	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		response := MessageStruct{"Error al enviar los datos del producto", 500}
		c.JSON(500,response)
		return
	}

	// cerramos el body
	defer resp.Body.Close()

	// leemos la respuesta de MeLi, que seran los datos del nuevo producto ya publicado
	data, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		response := MessageStruct{"Error al leer la respuesta de Mercado Libre", 500}
		c.JSON(500, response)
		return
	}
	response := MessageStruct{"Producto cargado con exito", 200}
	c.JSON(200, response)
}

// Obtener la lista de productos
func ItemList(user ReqUserData, chItemsList chan []itemData){

	idList, err := getItemsID(user)

	fmt.Println(user)
	fmt.Println(idList)
	fmt.Println(err)

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		chItemsList <- []itemData{}
		return
	}

	itemsList, err := getItemsList(idList)

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		chItemsList <- []itemData{}
		return
	}
	chItemsList <- itemsList
}

// obtener listado de preguntas
func QuestList (user ReqUserData, chQuestList chan []UnansweredQuest){

	resp, err := http.Get("https://api.mercadolibre.com/questions/search?seller_id=" +
		strconv.Itoa(user.IdMeli) + "&access_token=" + user.AccessToken +
		"&status=UNANSWERED&sort_fields=item_id,date_created&sort_type=ASC")

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err ))
		chQuestList <- []UnansweredQuest{}
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(fmt.Errorf("error %v", err.Error()))
		chQuestList <- []UnansweredQuest{}
		return
	}

	var respAux QuestsReq

	err = json.Unmarshal(data, &respAux)

	var questList []UnansweredQuest

	for i:=0; i<len(respAux.Questions); i++{
		aux := UnansweredQuest{
			"",
			respAux.Questions[i].ItemId,
			respAux.Questions[i].Id,
			respAux.Questions[i].Text,
		}
		questList = append(questList,aux)
	}

	chQuestList <- questList
}

// obtener listado de productos vendidos
func SoldList(user ReqUserData,chSoldList chan []soldItem){


	soldItemsList, err := getSoldItems(user)

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		chSoldList <- []soldItem{}
		return
	}
	var auxlist []soldItem

	for i:=0;i< len(soldItemsList);i++ {
		exist := false
		for j:=0;j< len(auxlist);j++{
			if soldItemsList[i].ItemId == auxlist[j].ItemId{
				exist = true
			}
		}
		if !exist {
			auxlist = append(auxlist,soldItemsList[i])
			for j:=i+1;j< len(soldItemsList);j++{
				if auxlist[len(auxlist)-1].ItemId == soldItemsList[j].ItemId{
					auxlist[len(auxlist)-1].Quantity = auxlist[len(auxlist)-1].Quantity + soldItemsList[j].Quantity
					auxlist[len(auxlist)-1].TotalPaidAmount = auxlist[len(auxlist)-1].TotalPaidAmount + soldItemsList[j].TotalPaidAmount
				}
			}
		}
	}

	chSoldList <- auxlist

}

// responder preguntas
func Answer(c*gin.Context){

	id := c.Query("id")

	user,err := testToken(id)

	if err != nil {
		fmt.Println(err)
		response := MessageStruct{"Error al tratar de obtener los datos del usuario", 401}
		c.JSON(401, response)
		return
	}


	// obtenemos los datos del formulario que recibimos
	question := c.Query("idq")
	answer :=  c.PostForm("answer")

	// creamos un string con los datos recibidos
	value := "{\n\"question_id\": " +  question + ", \n \"text\":\"" + answer + "\" \n}"

	// convertimos el string a un array de bytes
	b := []byte(value)

	// creamos la request con su respectivo header
	req, _ := http.NewRequest(http.MethodPost, "https://api.mercadolibre.com/answers",bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization","Bearer " + user.AccessToken)

	// realizamos la request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		response := MessageStruct{"Error al tratar de enviar la respuesta", 500}
		c.JSON(500, response)
	}

	// cerramos el body de la respuesta
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	datres := string(data)

	fmt.Println(datres)

	response := MessageStruct{"Respuesta enviada con exito", 200}
	c.JSON(200, response)
}

// envia lista de items, de preguntas, y de items vendidios
func Home(c*gin.Context) {

	id := c.Query("id")

	user,err := testToken(id)

	if err != nil {
		fmt.Println(err)
		response := MessageStruct{"Error al tratar de obtener los datos del usuario", 401}
		c.JSON(401, response)
		return
	}

	chSoldItemList := make(chan []soldItem)
	chItemsList := make(chan []itemData)
	chQuestList := make(chan []UnansweredQuest)

	go SoldList(user, chSoldItemList )
	go ItemList(user, chItemsList)
	go QuestList(user, chQuestList)

	soldItemList := <- chSoldItemList
	itemsList := <- chItemsList
	questList := <- chQuestList

	for i:=0; i<len(questList); i++{
		for j:=0; j<len(itemsList); j++{
			if questList[i].ItemId == itemsList[j].Id {
				questList[i].ItemTitle = itemsList[j].Title
			}
		}
	}

	homeResp := HomeStruct{soldItemList, itemsList, questList}


	c.JSON(200,homeResp)
}

func Export(c*gin.Context) {

	id := c.Query("id")

	user,err := testToken(id)

	if err != nil {
		fmt.Println(err)
		response := MessageStruct{"Error al tratar de obtener los datos del usuario", 401}
		c.JSON(401,response)
		return
	}

	chItemsList := make(chan []itemData)
	go ItemList(user, chItemsList)

	itemsList := <- chItemsList

	itemsListEx := ""

	for i:=0; i<len(itemsList); i++{

		aux := "('" + itemsList[i].Title + "'," +
				strconv.Itoa(itemsList[i].Quantity) +
				"," + fmt.Sprintf("%.2f",itemsList[i].Price) +
				"," + id + ")"

		if i==0{
			itemsListEx = aux
		}else{
			itemsListEx = itemsListEx + "," + aux
		}
	}

	itemsListEx = itemsListEx + ";"
	fmt.Println(itemsListEx)
	exportarProductos(itemsListEx)

	if err != nil {
		fmt.Println(err)
		response := MessageStruct{"Error al intentar exportar los datos", 500}
		c.JSON(500,response)
		return
	}

	response := MessageStruct{"Datos exportados con exito", 200}
	c.JSON(200, response)
}