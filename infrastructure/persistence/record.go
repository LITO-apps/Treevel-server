package persistence

import (
    "github.com/LITO-apps/Treevel-server/domain/models"
    "github.com/LITO-apps/Treevel-server/domain/repository"
    "github.com/gobuffalo/pop"
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
