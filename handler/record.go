package handler

import (
    "fmt"
    "github.com/LITO-apps/Treevel-server/usecase"
    "github.com/julienschmidt/httprouter"
    "net/http"
)

type RecordHandler interface {
    HandleGetAllRecords(http.ResponseWriter, *http.Request, httprouter.Params)
}

type recordHandler struct {
    recordUseCase usecase.RecordUseCase
}

func NewRecordHandler(ru usecase.RecordUseCase) RecordHandler {
    return &recordHandler{ru}
}

func (rh recordHandler) HandleGetAllRecords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
