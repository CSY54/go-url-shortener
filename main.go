package main

import app "github.com/CSY54/go-url-shortener/src"

func main() {
	r := app.Init(false)
	app.Run(r)
}
