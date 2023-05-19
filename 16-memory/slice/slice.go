package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	// == 既存の配列を参照するスライス
	b := a[:]
	fmt.Println(&b[0], len(b), cap(b))

	// 既存の配列の一部を参照するスライス
	c := a[1:3]
	fmt.Println(&c[0], len(c), cap(c))
	// => 0xc0000180f8 2 4
	// 大きさ(len)はスライスの要素数なので、2となる
	// 確保している要素数(cap)は
	// - 元の配列の要素数が5で、後ろの2要素はスライスの範囲外だが存在はしている
	// - またスライスの先頭が元の配列の[1]を指していて[0]を含まないため、元の配列より1小さい4になる

	// 何も参照していないスライス
	// d[0]は範囲外となるためエラー
	var d []int
	fmt.Println(len(d), cap(d))

	// == スライスと裏側の配列を同時に作成する場合

	// 初期の配列とリンクされているスライス
	// 0xc00001c0e0 4 4
	e := []int{1, 2, 3, 4}
	fmt.Println(&e[0], len(e), cap(e))

	e2 := append(e, 5)
	fmt.Println(&e2[0], len(e2), cap(e2))

	// サイズを持ったスライスを定義
	// 0xc00001c100 4 4
	f := make([]int, 4)
	fmt.Println(&f[0], len(f), cap(f))

	// サイズと容量を持ったスライスを定義
	// 0xc00001e300 4 7
	g := make([]int, 4, 7)
	fmt.Println(&g[0], len(g), cap(g))
}
