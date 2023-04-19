package main

import (
	"fmt"
	"os"
)

func main() {
	// 文字列のリストで全取得
	fmt.Println("# os.Environ()")
	environs := os.Environ()
	fmt.Println(len(environs))
	fmt.Println(environs[0])
	fmt.Println(environs)

	// 環境変数が埋め込まれた文字列を展開
	fmt.Println("# os.ExpandEnv()")
	fmt.Println(os.ExpandEnv("${HOME}/goblin"))

	// キーに対する値を取得 (有無をboolで返す)
	fmt.Println("# os.LookupEnv()")
	val, result := os.LookupEnv("DATABASE_HOST")
	fmt.Println("val:", val)
	fmt.Println("result ", result)

	// キーに対する値を取得
	fmt.Println("# os.GetEnv()")
	fmt.Println(os.Getenv("DATABASE_HOST"))
}
