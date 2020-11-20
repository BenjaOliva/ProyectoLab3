package server

import (
	"github.com/ProyectoLab3/pkg/tasks"
	"github.com/gin-gonic/gin"
)

func RunAPI() {

	//Se crea el servidor
	r := gin.Default()

	//listeners
	r.GET("/auth/code", tasks.GetToken)
	r.GET("/dashboard", tasks.GetItem)
	r.GET("/items", tasks.GetItemsOnly)

	//Puerto
	r.Run( ":8080")

	 /*
	http.HandleFunc("/hola", func(w http.ResponseWriter, peticion *http.Request) {
		io.WriteString(w, "Solicitaste hola")
	})
	direccion := ":8080" // Como cadena, no como entero; porque representa una dirección
	fmt.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))

	  */
}
