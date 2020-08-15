package persistence

import (
    "log"
    "time"

    "github.com/gobuffalo/pop"

    "github.com/LITO-apps/Treevel-server/domain/models"
    "github.com/LITO-apps/Treevel-server/domain/repository"
)

type playerPersistence struct {
    db *pop.Connection
}

func NewPlayerPersistence() repository.PlayerRepository {
    db, err := pop.Connect("development")
    if err != nil {
        // DB との接続ができない場合には強制終了
        log.Fatal(err)
    }
    log.Print("Succeeded in connecting database in `Player`")

    return &playerPersistence{db: db}
}

func (pp playerPersistence) GetAllPlayers() ([]models.Player, error) {
    var players []models.Player
    db := pp.db

    err := db.All(&players)
    if err != nil {
        return nil, err
    }

    return players, nil
}

func (pp playerPersistence) CreatePlayer(name string, t time.Time) error {
    player := models.Player{Name: name, LastLoginTime: t}
    db := pp.db

    _, err := db.ValidateAndCreate(&player)
    if err != nil {
        return err
    }

    return nil
}
