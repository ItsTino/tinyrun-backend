// cmd/server/main.go
package main

import (
	"tinyrun/internal/api"
	"tinyrun/internal/database"
)

func main() {
	db := database.SetupDB()
	router := api.SetupRouter(db)
	router.Run(":8080")
}
