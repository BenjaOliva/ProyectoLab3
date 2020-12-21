package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

//SUBESTRUCTURAS NECESARIAS
type Description struct {
	Plain_text string `json:"plain_text"`
}

type SaleTerms struct {
	Id         string `json:"id"`
	Value_name string `json:"value_name"`
}

type Picture struct {
	Source string `json:"source"`
}

type Atribute struct {
	Id         string `json:"id"`
	Value_name string `json:"value_name"`
}

//ESTRUCTURA QUE SE MANDA POR POST
type NewItem struct {
	Title              string      `json:"title"`
	CategoryId         string      `json:"category_id"`
	Price              float64     `json:"price"`
	Currency_id        string      `json:"currency_id"`
	Available_quantity int64       `json:"available_quantity"`
	Buying_mode        string      `json:"buying_mode"`
	Condition          string      `json:"condition"`
	Listing_type_id    string      `json:"listing_type_id"`
	Description        Description `json:"description"`
	Video_id           string      `json:"video_id"`
	Sale_terms         []SaleTerms `json:"sale_terms"`
	Pictures           []Picture   `json:"pictures"`
	Atributes          []Atribute  `json:"attributes"`
}

//ESTRUCTURA DE RESPUESTA
type Response struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type MyItem struct {
	Title      string `json:"title"`
	Quantity   string `json:"quantity"`
	Price      string `json:"price"`
	PictureUrl string `json:"url"`
	Condition  string `json:"condition"`
}

var ResponseNewItem MyItem

func NewProduct(c *gin.Context) {

	bodyFront, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Errorf("Error", err.Error())
		return
	}

	//itemToPost := string(bodyFront)
	//fmt.Println("Item del Form: ", itemToPost)

	json.Unmarshal(bodyFront, &ResponseNewItem)

	//En caso de no haber cargado una url en el formulario, se setea una url de imagen por defecto
	if ResponseNewItem.PictureUrl == "" {
		ResponseNewItem.PictureUrl = "http://mla-s2-p.mlstatic.com/968521-MLA20805195516_072016-O.jpg"
	}

	fmt.Println("Item a Publicar: ", ResponseNewItem)

	atributes := []Atribute{
		{
			Id:         "BRAND",
			Value_name: "Marca del producto",
		},
		{
			Id:         "EAN",
			Value_name: "7898095297749",
		},
	}

	pictures := []Picture{
		{
			Source: ResponseNewItem.PictureUrl,
		},
	}

	sale_terms := []SaleTerms{
		{
			Id:         "WARRANTY_TYPE",
			Value_name: "Garantía del vendedor",
		},
		{
			Id:         "WARRANTY_TIME",
			Value_name: "90 días",
		},
	}

	description := Description{
		Plain_text: "Item Publicado desde NetSpace\n",
	}

	//convertimos los parametros del body que enviamos desde el front a float y entero respectivamente porque vienen como strings
	price, err := strconv.ParseFloat(ResponseNewItem.Price, 64)
	quantity, err := strconv.ParseInt(ResponseNewItem.Quantity, 10, 64)

	newItem := NewItem{
		Title:              ResponseNewItem.Title,
		CategoryId:         "MLA3530",
		Price:              price,
		Currency_id:        "ARS",
		Available_quantity: quantity,
		Buying_mode:        "buy_it_now",
		Condition:          "new",
		Listing_type_id:    "gold_special",
		Description:        description,
		Video_id:           "YOUTUBE_ID_HERE",
		Sale_terms:         sale_terms,
		Pictures:           pictures,
		Atributes:          atributes,
	}

	jsonNewItem, _ := json.Marshal(newItem)

	//fmt.Println("JSONnewItem: ",string(jsonNewItem))

	//c.JSON(200, ResponseNewItem)

	responsePostNewItem, err := http.Post("https://api.mercadolibre.com/items?access_token="+TokenR.Access_token, "application/json; application/x-www-form-urlencoded", bytes.NewBuffer(jsonNewItem))

	if err != nil {
		fmt.Errorf("Error", err.Error())
		return
	}

	defer responsePostNewItem.Body.Close()

	response, err := ioutil.ReadAll(responsePostNewItem.Body)

	if err != nil {
		fmt.Errorf("Error", err.Error())
		return
	}

	bodyString := string(response)
	fmt.Println("Respuesta MELI: ", bodyString)

	json.Unmarshal(response, &ResponseNewItem)

	ResponseNewItem.PictureUrl = ""
	//c.JSON(200, ResponseNewItem)

}
