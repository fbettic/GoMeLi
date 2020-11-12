package api_back

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

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