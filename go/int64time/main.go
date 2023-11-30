package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// 何らかの処理を行う
	time.Sleep(1 * time.Microsecond)
	duration := time.Since(start)
	microseconds := int64(duration) / 1000
	fmt.Println(duration) //マイクロ秒
	fmt.Println(int64(duration)) //ナノ秒
	fmt.Println(microseconds) //マイクロ秒少数なし
}
