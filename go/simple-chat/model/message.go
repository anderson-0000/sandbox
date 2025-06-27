package model

type Message struct {
	Timestamp string `json:"timestamp"` // 送信日時
	Text      string `json:"text"`      // メッセージ本文
}
