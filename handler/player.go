package handler

import (
    "fmt"
    "net/http"

    "github.com/julienschmidt/httprouter"

    "github.com/LITO-apps/Treevel-server/usecase"
)

type PlayerHandler interface {
    HandleGetAllPlayers(http.ResponseWriter, *http.Request, httprouter.Params)
}

type playerHandler struct {
    playerUseCase usecase.PlayerUseCase
}

func NewPlayerHandler(pu usecase.PlayerUseCase) PlayerHandler {
    return &playerHandler{pu}
}

func (ph playerHandler) HandleGetAllPlayers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    players, err := ph.playerUseCase.GetAllPlayers()
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    _, err = fmt.Fprintln(w, players)
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
}
