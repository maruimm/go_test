package main

import (
	"bufio"
	"fmt"
	"strings"
)

// 示例：Peek、Read、Discard、Buffered
func main() {
	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	buf := bufio.NewReaderSize(sr, 0)
	b := make([]byte, 10)

	fmt.Println(buf.Buffered()) // 0
	s, _ := buf.Peek(5)
	s[0], s[1], s[2] = 'a', 'b', 'c'
	fmt.Printf("%d   %q \n", buf.Buffered(), s) // 16   "abcDE"

	buf.Discard(1)

	for n, err := 0, error(nil); err == nil; {
		n, err = buf.Read(b)
		fmt.Printf("%d %d  %q   %v\n", n,buf.Buffered(), b[:n], err)
	}
	// 5   "bcDEFGHIJK"   <nil>
	// 0   "LMNOP"   <nil>
	// 6   "QRSTUVWXYZ"   <nil>
	// 0   "123456"   <nil>
	// 0   "7890"   <nil>
	// 0   ""   EOF
}
