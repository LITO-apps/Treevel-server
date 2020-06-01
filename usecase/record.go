package usecase

import (
    "github.com/LITO-apps/Treevel-server/domain/models"
    "github.com/LITO-apps/Treevel-server/domain/repository"
)

type RecordUseCase interface {
    GetAllRecords() ([]models.Record, error)
    CreateRecord(int, int, bool, int, int, string) error
}

type recordUseCase struct {
    recordRepository repository.RecordRepository
}

func NewRecordUseCase(rr repository.RecordRepository) RecordUseCase {
    return &recordUseCase{rr}
}

func (ru recordUseCase) GetAllRecords() ([]models.Record, error) {
    records, err := ru.recordRepository.GetAllRecords()
    if err != nil {
        return nil, err
    }

    return records, nil
}

func (ru recordUseCase) CreateRecord(playerID int, stageID int, isClear bool, playTimes int, firstClearTimes int, minClearTime string) error {
    err := ru.recordRepository.CreateRecord(playerID, stageID, isClear, playTimes, firstClearTimes, minClearTime)
    if err != nil {
        return err
    }

    return nil
}
