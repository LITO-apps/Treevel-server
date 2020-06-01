package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "strconv"

    "github.com/LITO-apps/Treevel-server/models"
    "github.com/gobuffalo/nulls"
    "github.com/gobuffalo/pop"
    "github.com/julienschmidt/httprouter"

    "github.com/LITO-apps/Treevel-server/handler"
    "github.com/LITO-apps/Treevel-server/infrastructure/persistence"
    "github.com/LITO-apps/Treevel-server/usecase"
)

func main() {
    // player 関連の DI
    playerPersistence := persistence.NewPlayerPersistence()
    playerUseCase := usecase.NewPlayerUseCase(playerPersistence)
    playerHandler := handler.NewPlayerHandler(playerUseCase)

    // record 関連の DI
    recordPersistence := persistence.NewRecordPersistence()
    recordUseCase := usecase.NewRecordUseCase(recordPersistence)
    recordHandler := handler.NewRecordHandler(recordUseCase)

    // ルーティングの設定
    router := httprouter.New()
    router.GET("/get_all_players", playerHandler.HandleGetAllPlayers)
    router.GET("/get_all_records", recordHandler.HandleGetAllRecords)
    router.POST("/create_player", playerHandler.HandleCreatePlayer)
    router.POST("/create_record", createRecord)

    // サーバ起動
    fmt.Println("Server Start")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func createRecord(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    dump, err := httputil.DumpRequest(r, true)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    fmt.Println(string(dump))

    // parse post data
    playerID, err := strconv.Atoi(r.FormValue("player_id"))
    stageID, err := strconv.Atoi(r.FormValue("stage_id"))
    isClear, err := strconv.ParseBool(r.FormValue("is_clear"))
    playTimes, err := strconv.Atoi(r.FormValue("play_times"))
    firstClearTimes, err := strconv.Atoi(r.FormValue("first_clear_times"))
    minClearTime := r.FormValue("min_clear_time")

    db, err := pop.Connect("development")
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    record := models.Record {
        PlayerID: playerID, 
        StageId: stageID,
        IsClear: isClear,
        PlayTimes: playTimes,
        FirstClearTimes: nulls.NewInt(firstClearTimes),
        MinClearTime: nulls.NewString(minClearTime),
    }
    _, err = db.ValidateAndCreate(&record)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
}
