package main

import (
	"sql-db-con/Controller"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	Controller.NewAPIHandler(router)
	router.Run("localhost:8080")
}
