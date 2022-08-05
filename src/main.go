package main

import (
	"log"
	"os"
	"os/signal"
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

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server force shutdown...")
	close(quit)
	os.Exit(1)
}
