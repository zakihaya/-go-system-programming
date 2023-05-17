package main

import "fmt"

func main() {
	// リストを宣言
	var a [4]int
	fmt.Println(a)

	// リストを生成
	b := [4]int{}
	fmt.Println(b)

	//リストを生成（初期値つき）
	c := [4]int{0, 1, 2, 3}
	fmt.Println(c)

	// リストを生成（初期値つき/要素数は自動設定）
	d := [...]int{0, 1, 2}
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	fmt.Println(d[1])
}
