package main

import "fmt"

func main() {
	names := []string{"mike"}

	fmt.Println(names) // -> [mike]
	for _, n := range names {
		n = "bob"
		fmt.Println(n) // 実際はPrintlnではないがnを参照するコードがここのブロック内にある
	}
	fmt.Println(names) // -> [mike]

	// ポインタを確認すると異なることがわかる
	for i, v := range names {
		fmt.Printf("index: %d, original: %p, value derived from range: %p\n", i, &names[i], &v)
	}

	// インデックス経由なら元スライスも変更できる
	for i := range names {
		names[i] = "bob"
	}
	fmt.Println(names)
}
