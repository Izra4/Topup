package service

import (
	"TopUpWEb/entity"
	"TopUpWEb/repository"
)

type GameService interface {
	FindAll() ([]entity.Games, error)
	FindbyId(id uint) (entity.Games, error)
	GetListTopUpsByGameID(id uint) ([]entity.ListTopUp, error)
}

type gameService struct {
	gameRepo repository.GamesRepository
}

func NewGameService(gameRepo repository.GamesRepository) *gameService {
	return &gameService{gameRepo}
}

func (gs *gameService) FindAll() ([]entity.Games, error) {
	return gs.gameRepo.FindAll()
}

func (gs *gameService) FindbyId(id uint) (entity.Games, error) {
	return gs.gameRepo.FindbyId(id)
}

func (gs *gameService) GetListTopUpsByGameID(id uint) ([]entity.ListTopUp, error) {
	return gs.gameRepo.GetListTopUpsByGameID(id)
}
