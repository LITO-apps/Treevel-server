package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "strconv"
    "time"

    "github.com/LITO-apps/Treevel-server/models"
    "github.com/gobuffalo/nulls"
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

    db, err := pop.Connect("development")
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    var players []models.Player
    err = db.All(&players)
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

func getAllRecordsHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
    dump, err := httputil.DumpRequest(r, true)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    fmt.Println(string(dump))

    db, err := pop.Connect("development")
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    var records []models.Record
    err = db.All(&records)
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

func createPlayer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    dump, err := httputil.DumpRequest(r, true)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    fmt.Println(string(dump))

    // parse post data
    name := r.FormValue("name")

    t := time.Now()

    db, err := pop.Connect("development")
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    player := models.Player{Name: name, LastLoginTime: t}
    _, err = db.ValidateAndCreate(&player)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
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

func main() {
    // ルーティングの設定
    router := httprouter.New()
    router.GET("/", rootHandler)
    router.GET("/get_all_players", getAllPlayersHandler)
    router.GET("/get_all_records", getAllRecordsHandler)
    router.POST("/create_player", createPlayer)
    router.POST("/create_record", createRecord)

    // サーバ起動
    fmt.Println("Server Start")
    log.Fatal(http.ListenAndServe(":8080", router))
}
