package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gobuffalo/nulls"

	"github.com/LITO-apps/Treevel-server/usecase"
)

type RecordHandler interface {
    HandleGetAllRecords(http.ResponseWriter, *http.Request)
    HandleCreateRecord(http.ResponseWriter, *http.Request)
    HandleStageInfoGetAllUserMinClearTime(http.ResponseWriter, *http.Request)
    HandleStageInfoGetAvgClearRate(http.ResponseWriter, *http.Request)
}

type recordHandler struct {
    recordUseCase usecase.RecordUseCase
}

func NewRecordHandler(ru usecase.RecordUseCase) RecordHandler {
    return &recordHandler{ru}
}

func (rh recordHandler) HandleGetAllRecords(w http.ResponseWriter, r *http.Request) {
    records, err := rh.recordUseCase.GetAllRecords()
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    _, err = fmt.Fprintln(w, records)
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
}

func (rh recordHandler) HandleCreateRecord(w http.ResponseWriter, r *http.Request) {
    // parse post data
    playerID, err := strconv.Atoi(r.FormValue("player_id"))
    stageID, err := strconv.Atoi(r.FormValue("stage_id"))
    isClear, err := strconv.ParseBool(r.FormValue("is_clear"))
    playTimes, err := strconv.Atoi(r.FormValue("play_times"))
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    firstClearTimes, err := strconv.Atoi(r.FormValue("first_clear_times"))
    minClearTime, err := strconv.ParseFloat(r.FormValue("clear_time"), 32)

    err = rh.recordUseCase.CreateRecord(playerID, stageID, isClear, playTimes, nulls.Int{Int: firstClearTimes, Valid: err == nil}, nulls.NewFloat32(float32(minClearTime)))
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
}

func (rh recordHandler) HandleStageInfoGetAllUserMinClearTime(w http.ResponseWriter, r *http.Request) {
    // parse post data
    stageID, err := strconv.Atoi(r.FormValue("stage_id"))

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    ret, err := rh.recordUseCase.GetStageInfoAllUserMinClearTime(stageID)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    // nullじゃない場合
    if (ret.Valid) {
        _, err = fmt.Fprintln(w, ret.Float32)
    } else {
        _, err = fmt.Fprintln(w, nil)
    }

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
}

func (rh recordHandler) HandleStageInfoGetAvgClearRate(w http.ResponseWriter, r *http.Request) {
    // parse post data
    stageID, err := strconv.Atoi(r.FormValue("stage_id"))

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    ret, err := rh.recordUseCase.GetStageInfoAvgClearRate(stageID)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    _, err = fmt.Fprintln(w, ret)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
}
