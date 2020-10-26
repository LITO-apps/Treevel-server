package main

import (
    "fmt"
    "log"
    "net/http"

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
    router.POST("/create_record", recordHandler.HandleCreateRecord)
    router.POST("/stageinfo/get/all_user_min_clear_time", recordHandler.HandleStageInfoGetAllUserMinClearTime)

    // サーバ起動
    fmt.Println("Server Start")
    log.Fatal(http.ListenAndServe(":8080", router))
}
