package gostrings

import (
	"bytes"
	"fmt"
)

func GoReader() {
	x := []byte("你好，世界")

	r1 := bytes.NewReader(x)
	d1 := make([]byte, len(x))
	n, _ := r1.Read(d1)
	fmt.Println(r1.Size())
	fmt.Println(n, string(d1))
}
