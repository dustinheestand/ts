package store

import (
	"github.com/hustinheestand/ts/model"
	"github.com/jinzhu/gorm"
)

type GameStore struct {
	db *gorm.db
}

func NewGameStore(db *gorm.DB) *GameStore {
	return &GameStore{
		db: db,
	}
}

func (gs *GameStore) GetByID(int32 id) (*model.Game, error) {
	var g model.Game

	err := gs.db.Where(&model.Game{Id: id})
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &g, err
}

func (gs *GameStore) ListByUser(username string) ([]model.Game, error) {
	var (
		u     model.User
		games []model.Game
	)

	err := gs.db.Where(&model.User{Username: username}).First(&u).Error
	if err != nil {
		return nil, err
	}

	gs.db.Where(&model.Game{UserID: u.ID}).Find(&games)

	return games, nil
}
