package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, world!")
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	fmt.Println(portString)

}
