package main

import (
	"middleware/config"
	"middleware/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
