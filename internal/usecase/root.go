package usecase

import (
	"github.com/jmoiron/sqlx"
	"go-gateaway/internal/composities"
)

type UseCase struct {
	UserComposite *composities.UserComposite
}

func NewUseCase(db *sqlx.DB) *UseCase {
	userComposite := composities.NewUserComposite(db)

	return &UseCase{userComposite}
}
