package usecase

import (
    "github.com/LITO-apps/Treevel-server/domain/models"
    "github.com/LITO-apps/Treevel-server/domain/repository"
)

type PlayerUseCase interface {
    GetAllPlayers() ([]models.Player, error)
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
