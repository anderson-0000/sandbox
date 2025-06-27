package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-chat/model"
	"simple-chat/storage"
	"time"
)

// /messages へのリクエストを振り分けます
func MessagesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		getMessages(responseWriter, request)
	case http.MethodPost:
		postMessage(responseWriter, request)
	default:
		http.Error(responseWriter, "許可されていないメソッドです", http.StatusMethodNotAllowed)
	}
}

// 保存済みメッセージを返す
func getMessages(responseWriter http.ResponseWriter, request *http.Request) {
	messages, err := storage.ReadMessages()
	if err != nil {
		http.Error(responseWriter, "メッセージの読み込みに失敗しました", http.StatusInternalServerError)
		log.Printf("読み込みエラー: %v", err)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(responseWriter).Encode(messages)
	if err != nil {
		log.Printf("エンコードエラー: %v", err)
	}
}

// リクエストボディから受け取ったメッセージを保存
func postMessage(responseWriter http.ResponseWriter, request *http.Request) {
	var input struct {
		Text string `json:"text"`
	}
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil || input.Text == "" {
		http.Error(responseWriter, "無効なリクエストです", http.StatusBadRequest)
		return
	}
	msg := model.Message{
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Text:      input.Text,
	}
	err = storage.SaveMessage(msg)
	if err != nil {
		http.Error(responseWriter, "保存に失敗しました", http.StatusInternalServerError)
		log.Printf("保存エラー: %v", err)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseWriter.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(responseWriter).Encode(msg)
	if err != nil {
		log.Printf("JSON エンコードエラー: %v", err)
	}
}
