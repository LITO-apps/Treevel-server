package persistence

import (
    "github.com/gobuffalo/nulls"
    "github.com/gobuffalo/pop"
    "log"

    "github.com/LITO-apps/Treevel-server/domain/models"
    "github.com/LITO-apps/Treevel-server/domain/repository"
)

type recordPersistence struct {
    db *pop.Connection
}

func NewRecordPersistence() repository.RecordRepository {
    db, err := pop.Connect("development")
    if err != nil {
        // DB との接続ができない場合には強制終了
        log.Fatal(err)
    }
    log.Print("Succeed to connect database in `Record`")

    return &recordPersistence{db: db}
}

func (rp recordPersistence) GetAllRecords() ([]models.Record, error) {
    var records []models.Record
    db := rp.db

    err := db.All(&records)
    if err != nil {
        return nil, err
    }

    return records, nil
}

func (rp recordPersistence) CreateRecord(playerID int, stageID int, isClear bool, playTimes int, firstClearTimes nulls.Int, minClearTime nulls.String) error {
    record := models.Record {
        PlayerID: playerID,
        StageId: stageID,
        IsClear: isClear,
        PlayTimes: playTimes,
        FirstClearTimes: firstClearTimes,
        MinClearTime: minClearTime,
    }

    var db = rp.db

    _, err := db.ValidateAndCreate(&record)
    if err != nil {
        return err
    }

    return nil
}
