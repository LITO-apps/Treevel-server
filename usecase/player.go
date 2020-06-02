package usecase

import (
    "time"

    "github.com/LITO-apps/Treevel-server/domain/models"
    "github.com/LITO-apps/Treevel-server/domain/repository"
)

type PlayerUseCase interface {
    GetAllPlayers() ([]models.Player, error)
    CreatePlayer(string, time.Time) error
}

type playerUseCase struct {
    playerRepository repository.PlayerRepository
}

func NewPlayerUseCase(pr repository.PlayerRepository) PlayerUseCase {
    return &playerUseCase{pr}
}

func (pu playerUseCase) GetAllPlayers() ([]models.Player, error) {
    players, err := pu.playerRepository.GetAllPlayers()
    if err != nil {
        return nil, err
    }

    return players, nil
}

func (pu playerUseCase) CreatePlayer(name string, t time.Time) error {
    err := pu.playerRepository.CreatePlayer(name, t)
    if err != nil {
        return err
    }
    return nil
}
