package tasks

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

	fmt.Println("Item Guardado: " + ItemCaptured.Titulo)

	db, err := obtenerBaseDeDatos()
	if err != nil {
		fmt.Println("Error obteniendo base de datos: %v", err)
		return
	}
	// Terminar conexión al terminar función
	defer db.Close()

	// Ahora vemos si tenemos conexión
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error conectando: %v", err)
		return
	}
	// Listo, aquí ya podemos usar a db!
	fmt.Println("Conectado correctamente")

	insertar()

}

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "root"
	pass := ""
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "netspace"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func insertar() (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	// Preparamos para prevenir inyecciones SQL
	sentenciaPreparada, err := db.Prepare("INSERT INTO items (Titulo, Cantidad, Precio, Condicion, Usuario) VALUES(?, ?, ?,?,?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(ItemCaptured.Titulo, ItemCaptured.Cantidad, ItemCaptured.Precio, ItemCaptured.Condicion, UserDatasaved.User_Nickname)
	if err != nil {
		return err
	}
	return nil
}
