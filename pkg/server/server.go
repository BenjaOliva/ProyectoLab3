package server

import (
	"github.com/ProyectoLab3-master/pkg/tasks"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunAPI() {

	//Se crea el servidór
	r := gin.Default()

	//Cargamos los archivos que vamos a utilizar en la Plataforma
	r.LoadHTMLFiles("pkg/server/templates/index.html",
		"pkg/server/templates/signin.html",
		"pkg/server/templates/items.html")

	//EndPoints
	// - GetToken
	r.GET("/auth/code", tasks.GetToken, func(c *gin.Context) {
		// Renderizamos el HTML cargado previamente
		c.HTML(
			// seteamos status HTTP en 200 (OK)
			http.StatusOK,
			// Usamos el template cargado
			"index.html",
			// Pasamos los datos que queramos al archivo index, por ejemplo ID de Usuario en el titulo de pagina
			gin.H{
				"title": tasks.UserDatasaved.User_Nickname,
			},
		)
	})

	// - GetItem
	r.GET("/questions", tasks.GetAll)

	// - GetItemsOnly
	r.GET("/items", tasks.GetItemsOnly, func(c *gin.Context) {
		// Renderizamos el HTML cargado previamente
		c.HTML(
			// seteamos status HTTP en 200 (OK)
			http.StatusOK,
			// Usamos el template cargado
			"items.html",
			// Pasamos los datos que queramos al archivo index, por ejemplo ID de Usuario en el titulo de pagina
			gin.H{
				"title": tasks.UserDatasaved.User_Nickname,
			},
		)
	})

	// Sign-In de Plataforma
	r.GET("/ingresar", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"signin.html",
			// Pasamos los datos que queramos al archivo index, por ejemplo ID de Usuario en el titulo de pagina
			gin.H{
				"title": tasks.TokenR.User_id,
			},
		)
	})

	//Corremos el server en el puerto deseado
	r.Run(":8080")

}
