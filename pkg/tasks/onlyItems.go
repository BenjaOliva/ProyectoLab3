package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

//====================================================================================

type SoloItems struct{
	ItemsObtenidos [] Item
}

func GetItemsOnly(c *gin.Context) {

	// Creamos variable del tipo "Nuestros itemas" para almacenar los arrays de la informacion a mostrar en navegador, en formato JSON
	var OnlyItems SoloItems

	//Consulta para obtener todos los ITEMS de un usuario, pasamos 2 parametros, id de usuario generado por la funcion GetToken
	//y el parametro access token para validar y obtener los datos que solicitamos
	resp1, err := http.Get("https://api.mercadolibre.com/users/"+ strconv.Itoa(TokenR.User_id) +"/items/search?access_token=" + TokenR.Access_token)

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
	var items [] Item

	for i := 0; i < len(IdItemsVendedor.Id); i++ {
		//Get ´para obtener los datos de un item en concreto
		resp2, err := http.Get("https://api.mercadolibre.com/items/" + IdItemsVendedor.Id[i] + "?access_token=" + TokenR.Access_token)
		if err != nil {
			fmt.Errorf("Error",err.Error())
			return
		}

		dataItemsDetail, err := ioutil.ReadAll(resp2.Body)

		// Creamos una variable para manejar los datos Obtenidos y guardarlos en una variable del tipo ItemsMeli
		var item ItemMeli
		json.Unmarshal(dataItemsDetail, &item)

		//Variable temporal propia para almacenar los datos de la variable item previamente creada
		var itemTemp Item

		//Seteamos los datos de la variable del tipo MeLi en la propia
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
	OnlyItems.ItemsObtenidos = items

	c.JSON(200, OnlyItems)


}