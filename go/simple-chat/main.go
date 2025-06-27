package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-chat/handler"
)

func main() {
	http.HandleFunc("/messages", handler.MessagesHandler)
	port := ":8080"
	fmt.Println("チャットサーバーがポート", port, "で起動しました")
	log.Fatal(http.ListenAndServe(port, nil))
}
