package tasks

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//====================================================================================

var ItemCaptured Item

func SaveItem(c *gin.Context) {

	var ItemSelected string

	var position int

	//capturamos el item a guardar
	ItemSelected = c.Query("item")

	//Convertimos el numero obtenido de string a int
	position, _ = strconv.Atoi(ItemSelected)

	//Buscamos en el array el item en la posicion "position" y guardamos el Item en ItemCaptured
	ItemCaptured = OnlyItems.ItemsObtenidos[position]

	SaveOnDB(ItemCaptured)
}

func SaveOnDB(ItemCaptured Item) {

	fmt.Println("Item: " + ItemCaptured.Titulo)

}
