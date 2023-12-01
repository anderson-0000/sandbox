package main

import (
	"fmt"
	"net/http"
	"time"
)

func returnTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	unixTime := now.Unix()
	unixNano := now.UnixNano()
	fmt.Fprintf(w, "# TYPE current_time_seconds gauge\n")
	fmt.Fprintf(w, "current_time_seconds %v\n", unixTime)
	fmt.Fprintf(w, "# TYPE current_time_nanoseconds gauge\n")
	fmt.Fprintf(w, "current_time_nanoseconds %v\n", unixNano)
}

func main() {
	http.HandleFunc("/metric", returnTime)
//	if err := http.ListenAndServe(":50005", nil); err != nil {
//		panic(err)
//	}
// ↑のコメント行と↓は同じ処理
	err := http.ListenAndServe(":50005", nil)
	if err != nil {
		panic(err)
	}
}
