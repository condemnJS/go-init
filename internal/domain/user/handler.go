package user

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Handler struct {
	Service ServiceInterface
}

type HandlerInterface interface {
	Register(ctx fiber.Ctx) error
}

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{service}
}

func (handler *Handler) RegisterUserRoute(router *fiber.App) {
	router.Post("/register", handler.Register)
	//router.Post("/login", handlers.Login)
}

func (handler *Handler) Register(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"success": true,
		"message": "Successful created!",
	})
}
