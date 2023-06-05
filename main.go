package main

import (
	"go-echo/routes"
)

func main() {
	
	e := routes.Init()

	e.Logger.Fatal(e.Start("localhost:1234"))
}