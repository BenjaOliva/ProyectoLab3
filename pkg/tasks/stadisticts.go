package tasks

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var CantidadRegistros int
var CantidadUsers int
var NProducts string
var UserName string
var UserCount int
var UserName2 string
var UserCount2 int

func GetStats(c *gin.Context) {
	var name string = UserDatasaved.User_Nickname
	fmt.Println(name)
	obtener(name)
	obtenerUsers()
	newProducts()
	porcentajeDB(name)
	porcentajeDB2(name)
	//porcentaje()
}
func obtener(name string) (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		fmt.Println("Error obteniendo base de datos obtener: %v", err)
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
		fmt.Println("Error obteniendo base de datos obtenerUsers: %v", err)
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
		fmt.Println("Error obteniendo base de datos new products: %v", err)
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

func porcentajeDB(name string) {
	db, err := obtenerBaseDeDatos()

	fmt.Println(name)

	if err != nil {
		fmt.Println("Error obteniendo base de datos porcentajeDB: %v", err)
		return
	}
	err = db.QueryRow("SELECT `Usuario`,COUNT(`Titulo`) FROM `items` WHERE `Usuario` = ? ", name).Scan(&UserName, &UserCount)
	if err != nil {
		fmt.Println("Error al calcular el porcentaje de usuarios POR NAME: %v", err)
		return
	}
	fmt.Println("User1: ", UserCount, UserName)
	return

}

func porcentajeDB2(name string) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		fmt.Println("Error obteniendo base de datos porcentajeDB2: %v", err)
		return
	}
	err = db.QueryRow("SELECT `Usuario`, COUNT(`Titulo`) FROM `items` WHERE `Usuario` != ?", name).Scan(&UserName2, &UserCount2)
	if err != nil {
		fmt.Println("Error al calcular el porcentaje de usuarios: %v", err)
		return
	}
	fmt.Println("User2: ", UserCount2, UserName2)
	return

}

func porcentaje() {
	var rta1 int
	//rta1 = (UserCount + UserCount2) / UserCount
	var rta2 int
	rta2 = (UserCount2 + UserCount) / UserCount2
	UserCount = rta1
	UserCount2 = rta2
	fmt.Println(rta1, rta2)
	return
}
