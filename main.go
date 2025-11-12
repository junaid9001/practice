package main

import (
	"fmt"
	"os"

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
}
