package repository

import "github.com/LITO-apps/Treevel-server/domain/models"

type PlayerRepository interface {
    GetAllPlayers() ([]models.Player, error)
}
