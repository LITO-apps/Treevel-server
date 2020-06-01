package repository

import "github.com/LITO-apps/Treevel-server/domain/models"

type RecordRepository interface {
    GetAllRecords() ([]models.Record, error)
    CreateRecord(int, int, bool, int, int, string) error
}
