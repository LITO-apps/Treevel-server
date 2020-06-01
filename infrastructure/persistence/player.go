package persistence

import (
    "time"

    "github.com/gobuffalo/pop"

    "github.com/LITO-apps/Treevel-server/domain/models"
    "github.com/LITO-apps/Treevel-server/domain/repository"
)

type playerPersistence struct {}

func NewPlayerPersistence() repository.PlayerRepository {
    return &playerPersistence{}
}

func (pp playerPersistence) GetAllPlayers() ([]models.Player, error) {
    var players []models.Player

    db, err := pop.Connect("development")
    if err != nil {
        return nil, err
    }

    err = db.All(&players)
    if err != nil {
        return nil, err
    }

    return players, nil
}

func (pp playerPersistence) CreatePlayer(name string, t time.Time) error {
    player := models.Player{Name: name, LastLoginTime: t}
    
    db, err := pop.Connect("development")
    if err != nil {
        return err
    }

    _, err = db.ValidateAndCreate(&player)
    if (err != nil) {
        return err
    }
    
    return nil
}
