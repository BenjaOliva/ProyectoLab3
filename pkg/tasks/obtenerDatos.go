package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

//========================= STRUCTS PARA ITEMNS ====================
//Struct de imagenes para almacenar todas las imagenes del Item MeLi
type PictureMeli struct {
	Url string `json:"url"`
}

// Item Generado desde MeLi
type ItemMeli struct {
	Id                 string        `json:"id"`
	Title              string        `json:"title"`
	Price              float64       `json:"price"`
	Available_quantity int           `json:"available_quantity"`
	Pictures           []PictureMeli `json:"pictures"`
	Condition          string        `json:"condition"`
	Date               string        `json:"date_created"`
}

// --- STRUCT PRINCIPAL A MOSTRAR ---
type Item struct {
	Id              string
	Titulo          string
	Cantidad        int
	Precio          float64
	PrimeraImagen   string
	Condicion       string
	FechaDeCreacion string
}

//Struct para almacenar los ID de producto en un array de string
type ItemsIdMeli struct {
	Id []string `json:"results"`
}

//===================================================================

//========================= STRUCTS PARA PREGUNTAS SIN RESPONDER ====================

// Struct con estructura para almacenar lo que recibimos de MeLi en crudo
type QuestionMeli struct {
	Id           int          `json:"id"`
	Item_Id      string       `json:"item_id"`
	Date_created string       `json:"date_created"`
	Text         string       `json:"text"`
	Status       string       `json:"status"`
	From         QuestionFrom `json:"from"`
}

//Array de las preguntas del tipo de estructura mencionada anteriormente
type PreguntasSR struct {
	Questions []QuestionMeli `json:"questions"`
}

// --- STRUCT PRINCIPAL A MOSTRAR ---
type Unanswered_Question struct {
	Id            int
	Question_date string
	Title         string
	Question_text string
	FromUser      string
}

type QuestionFrom struct {
	id string `json:"id"`
}

//====================================================================================

//========================= STRUCTS PARA PRODUCTOS VENDIDOS ========================

type SingleItemMeli struct {
	Title string `json:"title"`
}

type Order_ItemsMeli struct {
	SingleItem      SingleItemMeli `json:"item"`
	Quantity        int            `json:"quantity"`
	Unit_price      float64        `json:"unit_price"`
	Full_Unit_Price float64        `json:"full_unit_price"`
}

type ResultMeli struct {
	Order_Items  []Order_ItemsMeli `json:"order_items"`
	Total_amount float64           `json:"total_amount"`
	Paid_amount  float64           `json:"paid_amount"`
	Date_closed  string            `json:"date_closed"`
}

type Sold_Item struct {
	Title         string
	Sold_Quantity int
	Unit_Price    float64
	Subtotal      float64
}

// --- STRUCT PRINCIPAL A MOSTRAR ---
type Sale_Order struct {
	Sold_Items     []Sold_Item
	Sale_date      string
	Total          float64
	Total_Delivery float64
}

type SoldItemMeli struct {
	Result []ResultMeli `json:"results"`
}

//====================================================================================

//========================= STRUCT PARA MOSTRAR LOS ARRAYS DE STRUCTS PRINCIPALES ========================

type OurItems struct {
	ItemsObtenidos       []Item
	Unanswered_Questions []Unanswered_Question
	Sales_Orders         []Sale_Order
}

//====================================================================================

var Questions []Unanswered_Question

// Creamos variable del tipo "Nuestros itemas" para almacenar los arrays de la informacion a mostrar en navegador, en formato JSON
var NuestrosItems OurItems

func GetAll(c *gin.Context) {

	//Consulta para obtener todos los ITEMS de un usuario, pasamos 2 parametros, id de usuario generado por la funcion GetToken
	//y el parametro access token para validar y obtener los datos que solicitamos
	resp1, err := http.Get("https://api.mercadolibre.com/users/" + strconv.Itoa(TokenR.User_id) + "/items/search?access_token=" + TokenR.Access_token)

	//Manejo de err
	if err != nil {
		fmt.Errorf("Error", err.Error())
		return
	}

	defer resp1.Body.Close()

	dataItemsId, err := ioutil.ReadAll(resp1.Body)

	var IdItemsVendedor ItemsIdMeli
	json.Unmarshal(dataItemsId, &IdItemsVendedor)

	//Array donde almacenamos los items convertidos
	var items []Item

	for i := 0; i < len(IdItemsVendedor.Id); i++ {
		//Get ´para obtener los datos de un item en concreto
		resp2, err := http.Get("https://api.mercadolibre.com/items/" + IdItemsVendedor.Id[i] + "?access_token=" + TokenR.Access_token)
		if err != nil {
			fmt.Errorf("Error", err.Error())
			return
		}

		dataItemsDetail, err := ioutil.ReadAll(resp2.Body)

		// Creamos una variable para manejar los datos Obtenidos y guardarlos en una variable del tipo ItemsMeli
		var item ItemMeli
		json.Unmarshal(dataItemsDetail, &item)

		//Variable temporal propia para almacenar los datos de la variable item previamente creada
		var itemTemp Item

		//Seteamos los datos de la variable del tipo MeLi en la propia

		itemTemp.Id = item.Id
		itemTemp.Titulo = item.Title
		itemTemp.Precio = item.Price
		itemTemp.PrimeraImagen = item.Pictures[0].Url
		itemTemp.Cantidad = item.Available_quantity
		itemTemp.Condicion = item.Condition
		itemTemp.FechaDeCreacion = item.Date
		//Guardamos en array
		items = append(items, itemTemp)

	}
	//Guardamos nuestros items en la variable Principal para mostrar
	NuestrosItems.ItemsObtenidos = items

	// Preguntas pendientes por responder por cada ítem ordenadas de las más antiguas a las más recientes.
	var Unanswered_Questions []Unanswered_Question

	for i := 0; i < len(IdItemsVendedor.Id); i++ {
		//Generamos la consulta de items pasando los parametros y ademas hacemos un sort por fecha de creacion del item, de forma asc
		//por lo cual mostrara las preguntas del item creado con mas antiguedad primero
		resp3, err := http.Get("https://api.mercadolibre.com/questions/search?item=" + IdItemsVendedor.Id[i] + "&access_token=" + TokenR.Access_token + "&sort_fields=date_created&sort_types=ASC")
		if err != nil {
			fmt.Errorf("Error", err.Error())
			return
		}

		dataQuestions, err := ioutil.ReadAll(resp3.Body)

		var questions PreguntasSR

		json.Unmarshal(dataQuestions, &questions)

		var UnansweredQuestiontemp Unanswered_Question

		for i := 0; i < len(questions.Questions); i++ {

			//Checkeamos que sean UNANSWERED para obtener solo aquellas preguntas no respondidas
			if len(questions.Questions) == 0 || questions.Questions[i].Status != "UNANSWERED" {
				continue
			}
			//Seteamos los datos de la variable del tipo MeLi en la propia
			UnansweredQuestiontemp.Id = questions.Questions[i].Id

			for j := 0; j < len(OnlyItems.ItemsObtenidos); j++ {
				if OnlyItems.ItemsObtenidos[j].Id == questions.Questions[i].Item_Id {
					UnansweredQuestiontemp.Title = OnlyItems.ItemsObtenidos[j].Titulo
				}
			}

			UnansweredQuestiontemp.Title = NuestrosItems.ItemsObtenidos[i].Titulo
			UnansweredQuestiontemp.Question_date = questions.Questions[i].Date_created
			UnansweredQuestiontemp.Question_text = questions.Questions[i].Text
			UnansweredQuestiontemp.FromUser = questions.Questions[i].From.id
			//Guardamos en array
			Unanswered_Questions = append(Unanswered_Questions, UnansweredQuestiontemp)
		}
	}

	NuestrosItems.Unanswered_Questions = Unanswered_Questions

	//  Ventas efectuadas
	resp2, err := http.Get("https://api.mercadolibre.com/orders/search?seller=" + strconv.Itoa(TokenR.User_id) + "&order.status=paid&access_token=" + TokenR.Access_token)

	defer resp2.Body.Close()

	dataSales, err := ioutil.ReadAll(resp2.Body)

	var soldItems SoldItemMeli
	json.Unmarshal(dataSales, &soldItems)

	var Sales_Orders []Sale_Order

	for i := 0; i < len(soldItems.Result); i++ {
		var Sale_Order_Temp Sale_Order
		Sale_Order_Temp.Sale_date = soldItems.Result[i].Date_closed
		Sale_Order_Temp.Total = soldItems.Result[i].Total_amount
		Sale_Order_Temp.Total_Delivery = soldItems.Result[i].Paid_amount
		for j := 0; j < len(soldItems.Result[i].Order_Items); j++ {
			var Sale_Order_Temp_Items Sold_Item
			Sale_Order_Temp_Items.Title = soldItems.Result[i].Order_Items[j].SingleItem.Title
			Sale_Order_Temp_Items.Unit_Price = soldItems.Result[i].Order_Items[j].Unit_price
			Sale_Order_Temp_Items.Sold_Quantity = soldItems.Result[i].Order_Items[j].Quantity
			Sale_Order_Temp_Items.Subtotal = soldItems.Result[i].Order_Items[j].Full_Unit_Price

			Sale_Order_Temp.Sold_Items = append(Sale_Order_Temp.Sold_Items, Sale_Order_Temp_Items)
		}
		Sales_Orders = append(Sales_Orders, Sale_Order_Temp)
	}

	NuestrosItems.Sales_Orders = Sales_Orders

	Questions = NuestrosItems.Unanswered_Questions

	//c.JSON(200, OurItems)

}
