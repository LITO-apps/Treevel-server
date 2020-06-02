package persistence

import (
    "github.com/gobuffalo/nulls"
    "github.com/gobuffalo/pop"

    "github.com/LITO-apps/Treevel-server/domain/models"
    "github.com/LITO-apps/Treevel-server/domain/repository"
)

type recordPersistence struct {}

func NewRecordPersistence() repository.RecordRepository {
    return &recordPersistence{}
}

func (rp recordPersistence) GetAllRecords() ([]models.Record, error) {
    var records []models.Record

    db, err := pop.Connect("development")
    if err != nil {
        return nil, err
    }

    err = db.All(&records)
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

    db, err := pop.Connect("development")
    if err != nil {
        return err
    }

    _, err = db.ValidateAndCreate(&record)
    if err != nil {
        return err
    }

    return nil
}
