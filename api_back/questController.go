package api_back

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

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