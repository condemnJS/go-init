package app

import (
	"go-gateaway/internal/config"
	"go-gateaway/internal/db"
	"go-gateaway/internal/router"
	"go-gateaway/internal/usecase"
)

type Application struct {
	DB      *db.DB
	Conf    *config.Config
	Router  *router.Router
	UseCase *usecase.UseCase
}

func Run() *Application {
	database := db.NewDB()
	conn := database.Init()

	conf := config.NewConfig()
	conf.Init()

	useCase := usecase.NewUseCase(conn)

	runner := router.NewRouter(useCase)
	runner.Run()

	return &Application{
		database,
		conf,
		runner,
		useCase,
	}
}
