package repository

import (
	"TopUpWEb/entity"
	"gorm.io/gorm"
)

type GamesRepository interface {
	FindAll() ([]entity.Games, error)
	FindbyId(id uint) (entity.Games, error)
	GetListTopUpsByGameID(id uint) ([]entity.ListTopUp, error)
}

type gamesRepository struct {
	db *gorm.DB
}

func NewGamesRepository(db *gorm.DB) *gamesRepository {
	return &gamesRepository{db}
}

func (g *gamesRepository) FindAll() ([]entity.Games, error) {
	var games []entity.Games
	if err := g.db.Find(&games).Error; err != nil {
		return nil, err
	}
	return games, nil
}

func (g *gamesRepository) FindbyId(id uint) (entity.Games, error) {
	var game entity.Games
	if err := g.db.First(&game, id).Error; err != nil {
		return entity.Games{}, err
	}
	return game, nil
}

func (g *gamesRepository) GetListTopUpsByGameID(id uint) ([]entity.ListTopUp, error) {
	var listTopUp []entity.ListTopUp
	if err := g.db.Where("games_id = ?", id).Find(&listTopUp).Error; err != nil {
		return nil, err
	}
	return listTopUp, nil
}
