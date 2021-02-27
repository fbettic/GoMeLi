package api_back

import (
	"bytes"
	"encoding/json"
	"errors"
	_ "errors"
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

	// realizamos el post de los datos a mercado libre
	resp, err := http.Post("https://api.mercadolibre.com/items?access_token=" + AccessToken,
		"application/json; application/x-www-form-urlencoded",
		bytes.NewBuffer(b))

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		return
	}

	// cerramos el body
	defer resp.Body.Close()

	// leemos la respuesta de MeLi, que seran los datos del nuevo producto ya publicado
	data, err := ioutil.ReadAll(resp.Body)

	fmt.Println(data)

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		c.String(http.StatusInternalServerError,"An error has occurred, please try again" )
		return
	}

	c.String(http.StatusOK, "Poducto cargado con exito")
}

// Obtener la lista de productos
func ItemList() ([]itemData,error){
	idList, err := getItemsID()

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		return []itemData{},errors.New("error al cargar los ID de productos")
	}

	itemsList, err := getItemsList(idList)

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		return []itemData{},errors.New("error al cargar los datos de productos")
	}
	return itemsList,nil
}

// obtener listado de preguntas
func QuestList () ([]UnansweredQuest,error){

	resp, err := http.Get("https://api.mercadolibre.com/questions/search?seller_id=" +
		strconv.Itoa(User.Id) + "&access_token=" + User.AccessToken +
		"&status=UNANSWERED&sort_fields=item_id,date_created&sort_type=ASC")

	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err ))
		return []UnansweredQuest{}, errors.New("error al pedir las preguntas")
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(fmt.Errorf("error %v", err.Error()))
		return []UnansweredQuest{}, errors.New("error al leer las preguntas")
	}

	var respAux QuestsReq

	err = json.Unmarshal(data, &respAux)

	var questList []UnansweredQuest

	for i:=0; i<len(respAux.Questions); i++{
		aux := UnansweredQuest{
			respAux.Questions[i].ItemId,
			respAux.Questions[i].Text,
		}
		questList = append(questList,aux)
	}

	return questList,nil
}

// obtener listado de productos vendidos
func SoldList() ([]soldItem,error){

	soldItemsList, err := getSoldItems()


	if err != nil {
		fmt.Println(fmt.Errorf("error %v",err.Error()))
		return []soldItem{},errors.New("error al cargar la lista de productos vendidos")
	}

	return soldItemsList,nil
}

// responder preguntas
func Answer(c*gin.Context){

	// obtenemos los datos del formulario que recibimos
	question := c.PostForm("question")
	answer :=  c.PostForm("answer")

	// creamos un string con los datos recibidos
	value := "{\n\"question_id\": " +  question + ", \n \"text\":\"" + answer + "\" \n}"

	// convertimos el string a un array de bytes
	b := []byte(value)

	// creamos la request con su respectivo header
	req, _ := http.NewRequest(http.MethodPost, "https://api.mercadolibre.com/answers",bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization","Bearer " + User.AccessToken)

	// realizamos la request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		println("error al obtener la respuesta")
	}

	// cerramos el body de la respuesta
	defer resp.Body.Close()

	// leemos la respuesta de MeLi, llegado a esta linea la respuesta ya fue posteada
	// esto es solo un mensaje de confirmacion de posteo
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		println("error al leer la respuesta")
	}

	println(string(data))
}

func Home(c*gin.Context) {
	soldItemList,err := SoldList()
	if err != nil {
		fmt.Println(err)
		return
	}

	itemsList,err := ItemList()
	if err != nil {
		fmt.Println(err)
		return
	}

	questList,err := QuestList()
	if err != nil {
		fmt.Println(err)
		return
	}

	homeResp := HomeStruct{soldItemList, itemsList, questList}

	c.JSON(200,homeResp)
}