package repository

import (
    "github.com/LITO-apps/Treevel-server/domain/models"
    "github.com/gobuffalo/nulls"
)

type RecordRepository interface {
    GetAllRecords() ([]models.Record, error)
    CreateRecord(int, string, bool, int, nulls.Int, nulls.Float32) error
    GetStageInfoAllUserMinClearTime(int) (nulls.Float32, error) 
    GetStageInfoAvgClearRate(int) (float32, error)
}
