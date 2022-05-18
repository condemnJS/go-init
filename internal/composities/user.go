package composities

import (
	"github.com/jmoiron/sqlx"
	"go-gateaway/internal/domain/user"
)

type UserComposite struct {
	user.RepositoryInterface
	user.ServiceInterface
}

func NewUserComposite(db *sqlx.DB) *UserComposite {
	rep := user.NewStorage(db)
	service := user.NewService(rep)

	return &UserComposite{
		rep,
		service,
	}
}
