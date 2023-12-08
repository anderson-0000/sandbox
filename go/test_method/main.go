package main

import "fmt"

type intType int

//メソッド
//func (レシーバの変数名 レシーバの型) メソッド名(パラメーターリスト) 戻り値の型 {
//メソッド名：println
func (value intType) println(){
	fmt.Println(value)
}

func main(){
	//intType型の変数intNumber1を宣言し、初期値として12345を代入
	var intNumber1 intType = 12345 //var版
	
	//:=版 12345をintType型に変換しintType型の値を作成、intNumber2という変数を宣言し、
	//初期値にintType型の値12345を代入
	intNumber2 := intType(12345)  

	//intType型のメソッドを呼び出す
	intNumber1.println()
        intNumber2.println()
	intType(12345).println() //intType型の値12345に対してprintlnメソッドを直接呼び出す
}
