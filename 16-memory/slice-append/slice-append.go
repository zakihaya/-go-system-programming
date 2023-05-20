package main

import "fmt"

func main() {
	// 長さ1, 確保された要素2のスライスを作成
	s := make([]int, 1, 2)
	fmt.Println(&s[0], len(s), cap(s))
	// => 0xc0000ac010 1 2

	// 1要素追加（確保された範囲内）
	s = append(s, 10)
	fmt.Println(&s[0], len(s), cap(s))
	// => 0xc000020120 2 2

	// さらに要素を追加（新しく配列が確保され直す）
	s = append(s, 20)
	fmt.Println(&s[0], len(s), cap(s))
	// => 0xc00001c0e0 3 4
}
