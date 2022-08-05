package main

import (
	"log"
	"taurus-backend/app"
)

func init() {
	log.Println("App init...")
	a := app.GetApp()
	a.Init()
	a.Check()
}

func main() {
	log.Println("App start...")
	app.GetApp().Run()
}
