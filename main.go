package main

import (
	"fmt"
	"os"
	"practice/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("error while loading env")
		return
	}
	name := os.Getenv("NAME")

	fmt.Print(name)

	router := gin.Default()

	routes.Authrouter(router)

	routes.Adminroute(router)

	router.Run(":" + os.Getenv("PORT"))
}
