package tasks

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

var CantidadRegistros int
var CantidadUsers int
var NProducts string

func GetStats(c *gin.Context) {
	var name string = UserDatasaved.User_Nickname
	fmt.Println(name)
	obtener(name)
	obtenerUsers()
	newProducts()
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
func obtenerUsers() {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		fmt.Println("Error obteniendo base de datos: %v", err)
		return
	}
	err = db.QueryRow("SELECT count(Usuario) FROM items").Scan(&CantidadUsers)
	if err != nil {
		fmt.Println("Error contando cantidad de usuarios: %v", err)
		return
	}
	fmt.Println(CantidadUsers)
	return
}

func newProducts() {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		fmt.Println("Error obteniendo base de datos: %v", err)
		return
	}
	var results *sql.Rows
	results, err = db.Query("SELECT COUNT(`Titulo`) FROM `items` WHERE `Condicion` = 'new' ")
	if err != nil {
		fmt.Println("Error contando: %v", err)
		return
	}
	for results.Next() {
		err = results.Scan(&NProducts)
		if err != nil {
			fmt.Println("Error al cargar array sexual")
		}
	}
	fmt.Println(NProducts)
	return
}
