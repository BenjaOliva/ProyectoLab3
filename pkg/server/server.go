package server

import (
	"net/http"

	"github.com/ProyectoLab3-master/pkg/tasks"
	"github.com/gin-gonic/gin"
)

func RunAPI() {

	//Se crea el servidór
	r := gin.Default()

	//Cargamos los archivos que vamos a utilizar en la Plataforma
	r.LoadHTMLGlob("pkg/server/templates/*")

	//EndPoints
	// - GetToken
	r.GET("/auth/code", tasks.GetToken, tasks.GetItemsOnly, tasks.GetAll, tasks.GetStats, func(c *gin.Context) {
		// Renderizamos el HTML cargado previamente
		c.HTML(
			// seteamos status HTTP en 200 (OK)
			http.StatusOK,
			// Usamos el template cargado
			"index.html",
			// Pasamos los datos que queramos al archivo index, por ejemplo "nickname" de Usuario en el titulo de pagina
			gin.H{
				"title":            tasks.UserDatasaved.User_Nickname,
				"itemCounter":      len(tasks.OnlyItems.ItemsObtenidos),
				"PendingQuestions": len(tasks.NuestrosItems.Unanswered_Questions),
				"soldItemsCounter": len(tasks.NuestrosItems.Sales_Orders),
				"userName1":        tasks.UserName,
				"userCount1":       tasks.UserCount,
				"userName2":        tasks.UserName2,
				"userCount2":       tasks.UserCount2,
			},
		)
	})

	// - Get Questions
	r.GET("/questions", tasks.GetAll, func(c *gin.Context) {
		// Renderizamos el HTML cargado previamente
		c.HTML(
			// seteamos status HTTP en 200 (OK)
			http.StatusOK,
			// Usamos el template cargado
			"questions.html",
			// Pasamos los datos que queramos al archivo index, por ejemplo ID de Usuario en el titulo de pagina
			gin.H{
				"title":     tasks.UserDatasaved.User_Nickname,
				"Preguntas": tasks.Questions,
			},
		)
	})

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
				"title":     tasks.UserDatasaved.User_Nickname,
				"Productos": tasks.OnlyItems,
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
				"title":     tasks.TokenR.User_id,
				"Productos": tasks.OnlyItems,
			},
		)
	})

	r.GET("/crear", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"newProduct.html",
			// Pasamos los datos que queramos al archivo index, por ejemplo ID de Usuario en el titulo de pagina
			gin.H{
				"title": tasks.UserDatasaved.User_Nickname,
			},
		)
	})

	r.POST("/submit", tasks.NewProduct)

	r.POST("/answer", tasks.AnswerQuestion)

	r.GET("/stats", tasks.GetStats, func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"stats.html",
			// Pasamos los datos que queramos al archivo index, por ejemplo ID de Usuario en el titulo de pagina
			gin.H{
				"title":         tasks.UserDatasaved.User_Nickname,
				"Contador":      tasks.CantidadRegistros,
				"contadorUsers": tasks.CantidadUsers,
				"NewProducts":   tasks.NProducts,
			},
		)
	})

	//Corremos el server en el puerto deseado
	r.Run(":8080")

}
