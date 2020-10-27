package persistence

import (
    "fmt"
    "log"

    "github.com/gobuffalo/nulls"
    "github.com/gobuffalo/pop"

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
    log.Print("Succeeded in connecting database in `Record`")

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

func (rp recordPersistence) CreateRecord(playerID int, stageID int, isClear bool, playTimes int, firstClearTimes nulls.Int, minClearTime nulls.Float32) error {
    record := models.Record {
        PlayerID: playerID,
        StageId: stageID,
        IsClear: isClear,
        PlayTimes: playTimes,
        FirstClearTimes: firstClearTimes,
        ClearTime: minClearTime,
    }

    var db = rp.db

    _, err := db.ValidateAndCreate(&record)
    if err != nil {
        return err
    }

    return nil
}

func (rp recordPersistence) GetStageInfoAllUserMinClearTime(stageID int) (nulls.Float32, error) {
    db := rp.db
    record := []models.Record{}

    query := db.Where("stage_id = ? AND clear_time IS NOT NULL", stageID).Order("clear_time asc")
    err := query.All(&record)

    if (err != nil) {
        return nulls.Float32{}, err
    }

    if (len(record) > 0) {
        // ソートしているので0が一番小さいやつ
        min := record[0]

        if (min.ClearTime.Valid) {
            return nulls.Float32(min.ClearTime) , nil
        }
    }
    return nulls.Float32{}, nil
}

func (rp recordPersistence) GetStageInfoAvgClearRate(stageID int) (float32, error) {
    db := rp.db
    clearRecords := []models.Record{}

    // クリアしたプレイヤー数を取得
    err := db.Where("stage_id = ?", stageID).Where("is_clear = ?", 1).All(&clearRecords)
    clearNum := len(clearRecords)

    if (err != nil) {
        return 0, err
    }

    // プレイしたプレイヤー数を取得
    allRecords := []models.Record{}
    err = db.Where("stage_id = ?", stageID).All(&allRecords)
    playNum := len(allRecords)

    fmt.Println(playNum)

    // 分母が不正の場合
    if (playNum <= 0) {
        return 0, nil
    }

    return float32(clearNum) / float32(playNum), nil
}
