package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in, in1, in2, in3 string
	var n int
	var err error

	fmt.Print("continue? (Y/n) >")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		in = scanner.Text()
		fmt.Println("in: ", in)
		switch in {
		case "", "Y":
			fmt.Println("デフォルトの処理をします")
			goto L
		case "n":
			fmt.Println("何もしない")
			goto L
		default:
			fmt.Println("コマンドが不正なのでもう一度入力を促す")
			continue
		}
	}
L:

	// https://golang.org/pkg/fmt/#Scan
	fmt.Print("fmt.Scan >")
	n, err = fmt.Scan(&in1, &in2, &in3) // 3つの入力値を受け付ける
	if err != nil {
		panic(err)
	}
	fmt.Println("in1: ", in1)
	fmt.Println("in2: ", in2)
	fmt.Println("in3: ", in3)
	fmt.Println("n: ", n)

	// https://golang.org/pkg/fmt/#Scanln
	fmt.Print("fmt.Scanln >")
	n, err = fmt.Scanln(&in1, &in2, &in3) // 3つの入力値を受け付ける
	if err != nil {
		panic(err)
	}
	fmt.Println("in1: ", in1)
	fmt.Println("in2: ", in2)
	fmt.Println("in3: ", in3)
	fmt.Println("n: ", n)

}
