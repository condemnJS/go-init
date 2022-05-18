package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type ServiceInterface interface {
	Register(ctx *fiber.Ctx)
	Login(ctx *fiber.Ctx)
}

type Service struct {
	*Storage
}

func NewService(storage *Storage) *Service {
	return &Service{storage}
}

func (usrService *Service) Register(ctx *fiber.Ctx) {
	usrService.Storage.Create()

	fmt.Println("CREATE USER")
}

func (usrService *Service) Login(ctx *fiber.Ctx) {

}
