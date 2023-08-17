package main

import (
    "fmt"
)

type Shape interface {  // Shapeという名前のインターフェース
    Area() float64  // Aera()はメソッドでfloat64を返す
}

type Rectangle struct { // Rectangleという名前の構造体
    Width  float64
    Height float64
}

type Circle struct {  // Circleという名前の構造体
    Radius float64
}

//メソッドはオブジェクトが持つ処理や手続きの方法をが書いてある
func (r Rectangle) Area() float64 { // RectangleにAreaという名前のメソッドを追加した
    return r.Width * r.Height
}

func (c Circle) Area() float64 { // Circleという構造体にAreaという名前のメソッドを追加した
    return 3.14 * c.Radius * c.Radius
}

func main() {
    var s Shape  // Shape型の変数sを定義
    s = Rectangle{Width: 3, Height: 4}  // Rectanble構造体のインスタンスを作成しsに代入
    fmt.Println(s)  // {3 4}
    fmt.Println(s.Area()) // 12  s.Area()が呼び出されることで面積が計算される。
    fmt.Println(fmt.Sprint("四角の面積: ", s.Area()))

    s = Circle{Radius: 5}
    fmt.Println(fmt.Sprint("円の面積: ", s.Area()))
}
