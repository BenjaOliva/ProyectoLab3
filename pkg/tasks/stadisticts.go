package tasks

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var CantidadRegistros int

func GetStats(c *gin.Context) {
	var name string = UserDatasaved.User_Nickname
	fmt.Println(name)
	obtener(name)
}
func obtener(name string) (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		fmt.Println("Error obteniendo base de datos: %v", err)
		return
	}
	err = db.QueryRow("SELECT count(*) FROM `items` WHERE `Usuario` = ? ", name).Scan(&CantidadRegistros)
	if err != nil {
		fmt.Println("Error contando: %v", err)
		return
	}
	fmt.Println(CantidadRegistros)
	return
}
