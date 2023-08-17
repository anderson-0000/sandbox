package main

import (
    "fmt"
    "sort"
    "strconv"
    "strings"
)

func main() {
    slice := []string{
        "/test/version:v1.2.20",
        "/test/version:v1.2.21",
        "/test/version:v1.2.3",
        "/test/version:v1.2.3",
        "/test/version:v1.2.3.0",
        "/test/version:v1.2.3.2",
        "/test/version:v1.2.5.3",
        "/test/version:v1.2.0",
        "/test/version:v1.2.30",
        "/test/version:v1.2.2",
        "/test/version:v2.3.1",
        "/test/version:v2.4.3",
        "/test/version:v2.30.3",
        "/test/version:v2.3.12",
        "/test/version:v20.3.12",
        "/test/version:v3.3.12",
    }
    
    var versions []string
    for _, value := range slice {
	versions = append(versions, strings.Split(value, ":")[1])
    }
 
    sort.Slice(versions, func(a, b int) bool {  //versionsスライスをソートする(aとbを比較)
        return compareVersions(versions[a], versions[b])
    })
    
    fmt.Println(versions)
}

func compareVersions(a, b string) bool {
    aParts := strings.Split(a[1:], ".") //一文字目より後ろの値を.で分割してスライスにする
    bParts := strings.Split(b[1:], ".")
    
    for i := 0; i < len(aParts) && i < len(bParts); i++ {
        aInt, _ := strconv.Atoi(aParts[i]) //文字列を数値に変換
        bInt, _ := strconv.Atoi(bParts[i])
        
        if aInt < bInt {
            return true
        } else if aInt > bInt {
            return false
        }
    }
    
    return len(aParts) < len(bParts) //aとbが同じときやバッチバージョンの下にビルドメタデータがついている場合など
}

