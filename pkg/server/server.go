package server

import (
	"github.com/ProyectoLab3-master/pkg/tasks"
	"github.com/gin-gonic/gin"
)

func RunAPI() {

	//Se crea el servidor
	r := gin.Default()

	//EndPoints
	// - GetToken
	r.GET("/auth/code", tasks.GetToken)

	// - GetItem
	r.GET("/dashboard", tasks.GetAll)

	// - GetItemsOnly
	r.GET("/items", tasks.GetItemsOnly)

	//Puerto
	r.Run( ":8080")

}
