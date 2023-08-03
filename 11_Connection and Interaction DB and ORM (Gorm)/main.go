package main

import (
	"gorm/config"
	"gorm/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
