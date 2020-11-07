package usecase

import (
    "github.com/LITO-apps/Treevel-server/domain/models"
    "github.com/LITO-apps/Treevel-server/domain/repository"
    "github.com/gobuffalo/nulls"
)

type RecordUseCase interface {
    GetAllRecords() ([]models.Record, error)
    CreateRecord(int, string, bool, int, nulls.Int, nulls.Float32) error
    GetStageInfoAllUserMinClearTime(int) (nulls.Float32, error) 
    GetStageInfoAvgClearRate(int) (float32, error)
}

type recordUseCase struct {
    recordRepository repository.RecordRepository
}

func NewRecordUseCase(rr repository.RecordRepository) RecordUseCase {
    return &recordUseCase{rr}
}

func (ru recordUseCase) GetStageInfoAllUserMinClearTime (stageID int) (nulls.Float32, error) {
    clearTime , err := ru.recordRepository.GetStageInfoAllUserMinClearTime(stageID)

    if (err != nil) {
        return nulls.Float32{}, err
    }

    return clearTime, nil
}

func (ru recordUseCase) GetStageInfoAvgClearRate (stageID int) (float32, error) {
    clearTime , err := ru.recordRepository.GetStageInfoAvgClearRate(stageID)

    if (err != nil) {
        return 0, err
    }

    return clearTime, nil
}

func (ru recordUseCase) GetAllRecords() ([]models.Record, error) {
    records, err := ru.recordRepository.GetAllRecords()
    if err != nil {
        return nil, err
    }

    return records, nil
}

func (ru recordUseCase) CreateRecord(playerID int, stageID string, isClear bool, playTimes int, firstClearTimes nulls.Int, clearTime nulls.Float32) error {
    err := ru.recordRepository.CreateRecord(playerID, stageID, isClear, playTimes, firstClearTimes, clearTime)
    if err != nil {
        return err
    }

    return nil
}
