package main

import (
    "fmt"
    "github.com/LITO-apps/Treevel-server/handler"
    "github.com/LITO-apps/Treevel-server/infrastructure/persistence"
    "github.com/LITO-apps/Treevel-server/usecase"
    "github.com/julienschmidt/httprouter"
    "log"
    "net/http"
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

    // サーバ起動
    fmt.Println("Server Start")
    log.Fatal(http.ListenAndServe(":8080", router))
}
