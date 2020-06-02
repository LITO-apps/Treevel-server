package repository

import (
    "time"

    "github.com/LITO-apps/Treevel-server/domain/models"
)

type PlayerRepository interface {
    GetAllPlayers() ([]models.Player, error)
    CreatePlayer(string, time.Time) error
}
