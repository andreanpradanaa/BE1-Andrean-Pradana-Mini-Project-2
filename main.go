package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniproject2/modules/admin"
	user "miniproject2/modules/customer"
	"miniproject2/utils/db"
)

func main() {
	router := gin.New()

	// open connection db
	dbCrud := db.GormMysql()

	fmt.Println("database connected..!")

	versionRoute := router.Group("/v1")

	adminHandler := admin.NewRouter(dbCrud)
	adminHandler.Handle(versionRoute)

	userHandler := user.NewRouter(dbCrud)
	userHandler.Handle(versionRoute)

	errRouter := router.Run(":8080")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}

}
