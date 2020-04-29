package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/julienschmidt/httprouter"
)

func rootHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params)  {
	dump, err := httputil.DumpRequest(r, true)

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(dump))

	// FIXME: Unhandled error
	fmt.Fprint(w, "hello world!\n")
}

func main()  {
	// ルーティングの設定
	router := httprouter.New()
	router.GET("/", rootHandler)

	// サーバ起動
	fmt.Println("Server Start")
	log.Fatal(http.ListenAndServe(":8080", router))
}
