package game

import (
	"github.com/dustinheestand/ts/model"
)

type Store interface {
	GetByID(int32) (*model.Game, error)
	ListByUser(username string) ([]model.Game, error)
	UpdateGame(*model.Article, *model.Game) error
}
