package router

import (
	"github.com/gofiber/fiber/v2"
	"go-gateaway/internal/domain/user"
	"go-gateaway/internal/usecase"
	"log"
	"os"
)

type Router struct {
	cfg      *Config
	handlers *Handlers
}

type Config struct {
	Port string
}

type Handlers struct {
	user.Handler
}

func NewConfig() *Config {
	return &Config{Port: os.Getenv("APP_LISTEN_PORT")}
}

func NewHandlers(useCase *usecase.UseCase) *Handlers {
	userHandler := user.NewHandler(useCase.UserComposite.ServiceInterface)

	return &Handlers{*userHandler}
}

func NewRouter(useCase *usecase.UseCase) *Router {
	cfg := NewConfig()
	handlers := NewHandlers(useCase)

	return &Router{cfg, handlers}
}

func (router *Router) Run() {
	app := fiber.New()

	router.handlers.RegisterUserRoute(app)

	err := app.Listen(":" + router.cfg.Port)

	if err != nil {
		log.Fatalf("Fiber Router Error: %s", err)
	}
}
