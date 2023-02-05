package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var csvSource = `1,100  ,a,あ,x
2,100  ,b,い,x
3,101  ,c,う,x
4,102  ,d,え,x
5,102  ,e,お,x`

func main() {
	reader := strings.NewReader(csvSource)
	csvReader := csv.NewReader(reader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[0], line[1], line[2:4])
	}
}
