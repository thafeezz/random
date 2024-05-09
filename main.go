package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thafeezz/random/helpers"
)

func main() {
	router := gin.Default()

	db, err := helpers.ConnectDb()
	if err != nil {
		panic("Error connecting to database")
	}

	handler := helpers.NewHandler(db)

	router.GET("/emails", handler.GetEmails)
	router.POST("/addEmail", handler.AddEmail)

	// defer function call til before main returns
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic("Failed to close database")
		}
	}(db) // call lambda with db

	address := "localhost:8000"
	err = router.Run(address)
	if err != nil {
		fmt.Println("Failed to run server at address " + address)
		return
	}
}
