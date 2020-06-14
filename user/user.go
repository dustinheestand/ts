package user

import (
	"github.com/dustinheestand/ts/model"
)

type Store interface {
	GetByEmail(string) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
}
