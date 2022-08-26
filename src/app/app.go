package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
	"taurus-backend/api"
	"taurus-backend/constant"
	"taurus-backend/db"
)

var (
	once sync.Once
	app  *App
)

type App struct {
	srv *api.Server
	env *constant.Env
	db  *sql.DB
}

func GetApp() *App {
	once.Do(func() {
		a := &App{}
		app = a
	})
	return app
}

func (a *App) Check() {
	log.Println("App checking...")

	if err := a.db.Ping(); err != nil {
		log.Fatal("open database fail")
	}
	log.Println("database connect success")
}

func (a *App) Init() {
	log.Println("App init...")
	e := &constant.Env{}
	e.Init()
	a.env = e

	if "prod" == e.Stage {
		gin.SetMode(gin.ReleaseMode)
	}

	srv := api.NewServer()
	srv.Init()
	a.srv = srv

	dbHandler, err := db.Init(e)
	if err != nil {
		log.Fatal("db open fail", err)
	}
	a.db = dbHandler
}

func (a *App) Run() {
	log.Println("App run...")
	a.srv.Run()
}
