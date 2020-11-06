package handler

import (
    "fmt"
    "net/http"
    "time"

    "github.com/LITO-apps/Treevel-server/usecase"
)

type PlayerHandler interface {
    HandleGetAllPlayers(http.ResponseWriter, *http.Request)
    HandleCreatePlayer(http.ResponseWriter, *http.Request)
}

type playerHandler struct {
    playerUseCase usecase.PlayerUseCase
}

func NewPlayerHandler(pu usecase.PlayerUseCase) PlayerHandler {
    return &playerHandler{pu}
}

func (ph playerHandler) HandleGetAllPlayers(w http.ResponseWriter, r *http.Request) {
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

func (ph playerHandler) HandleCreatePlayer(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    t := time.Now()

    err := ph.playerUseCase.CreatePlayer(name, t)
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
}
