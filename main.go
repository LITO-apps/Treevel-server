package main

import (
    "fmt"
    "github.com/LITO-apps/Treevel-server/models"
    "log"
    "net/http"
    "net/http/httputil"

    "github.com/gobuffalo/pop"
    "github.com/julienschmidt/httprouter"
)

func rootHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
    dump, err := httputil.DumpRequest(r, true)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    fmt.Println(string(dump))

    _, err = fmt.Fprint(w, "hello world!\n")

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
}

func getAllPlayersHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
    dump, err := httputil.DumpRequest(r, true)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    fmt.Println(string(dump))

    tx, err := pop.Connect("development")
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    var players []models.Player
    err = tx.All(&players)
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    _, err = fmt.Fprint(w, players)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
}

func getAllRecordsHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
    dump, err := httputil.DumpRequest(r, true)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    fmt.Println(string(dump))

    tx, err := pop.Connect("development")
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    var records []models.Record
    err = tx.All(&records)
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    _, err = fmt.Fprint(w, records)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
}

func main()  {
    // ルーティングの設定
    router := httprouter.New()
    router.GET("/", rootHandler)
    router.GET("/get_all_players", getAllPlayersHandler)
    router.GET("/get_all_records", getAllRecordsHandler)

    // サーバ起動
    fmt.Println("Server Start")
    log.Fatal(http.ListenAndServe(":8080", router))
}
