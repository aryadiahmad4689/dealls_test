package app

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	appdb "github.com/aryadiahmad4689/dealls_test/src/app/db/sqlite"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jabardigitalservice/golog/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Debug bool
}
type App struct {
	Logger *logger.Logger
	router *chi.Mux

	db *Db
}

type Db struct {
	Master *sql.DB
	Slave  *sql.DB
}

func Init() *App {
	log := logger.Init()

	return initProd(log)
}

func initProd(logg *logger.Logger) *App {

	var (
		ctx = context.Background()

		app = &App{
			Logger: logg,
			router: InitChi(Config{
				Debug: false,
			}),
			db: &Db{
				Master: appdb.InitSqliteMaster(ctx),
				Slave:  appdb.InitSqliteSlave(ctx),
			},
		}
	)

	app.db.Master.SetMaxIdleConns(viper.GetInt("DB_MAX_IDLE_CONNS_MASTER"))
	app.db.Master.SetMaxOpenConns(viper.GetInt("DB_MAX_OPEN_CONNS_MASTER"))

	app.db.Slave.SetMaxIdleConns(viper.GetInt("DB_MAX_IDLE_CONNS_SLAVE"))
	app.db.Slave.SetMaxOpenConns(viper.GetInt("DB_MAX_OPEN_CONNS_SLAVE"))

	return app
}

func InitChi(config Config) *chi.Mux {
	var router = chi.NewRouter()

	if config.Debug {
		router.Use(middleware.Logger)
	}

	return router
}

func (app *App) Run() error {

	var port = os.Getenv("APP_PORT")
	var host = "0.0.0.0:" + port
	app.Logger.Info(&logger.LoggerData{
		Category: logger.LoggerApp,
		Service:  "dating_apps_dealls",
		Method:   "startup",
		Version:  os.Getenv("APP_VERSION"),
		AdditionalInfo: map[string]interface{}{
			"host": host,
		},
	}, "running")
	return http.ListenAndServe(host, app.router)
}

func (app *App) GetDb() *Db {
	return app.db
}

func (app *App) SetDb(db *Db) {
	app.db = db
}

func (app *App) GetHttpRouter() *chi.Mux {
	return app.router
}
