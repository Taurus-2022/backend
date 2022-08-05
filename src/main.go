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
	taurus := app.GetApp()
	log.Println("App start...")
	taurus.Run()
	log.Println("App start handle async task...")
	taurus.HandleAsyncTask()
}
