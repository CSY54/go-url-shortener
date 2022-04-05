package main

import (
	"os"

	app "github.com/CSY54/go-url-shortener/src"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := app.Init(false)
	app.Run(r, os.Getenv("ADDR"))
}
