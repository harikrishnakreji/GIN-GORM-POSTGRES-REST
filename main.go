package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harikrishnakreji/GO-GORM-REST/config"
	"github.com/harikrishnakreji/GO-GORM-REST/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRouter(router)

	router.Run(":8080")
}
