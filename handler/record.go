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
    minClearTime := r.FormValue("min_clear_time")

    err = rh.recordUseCase.CreateRecord(playerID, stageID, isClear, playTimes, nulls.Int{Int: firstClearTimes, Valid: err == nil}, nulls.NewString(minClearTime))
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
}
