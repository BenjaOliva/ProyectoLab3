package server

import (
	"github.com/ProyectoLab3/pkg/tasks"
	"github.com/gin-gonic/gin"
)

func RunAPI() {

	r := gin.Default()

	r.GET("/auth/code", tasks.GetToken)
	//r.GET("/dashboard", controller.GetDashboard)
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
