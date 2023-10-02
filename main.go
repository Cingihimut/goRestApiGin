package main

import (
	"github.com/Cingihimut/goRestApiGin.git/src/config"
	"github.com/Cingihimut/goRestApiGin.git/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.DB()

	routes.Api(r, db)
	r.Run()
}
