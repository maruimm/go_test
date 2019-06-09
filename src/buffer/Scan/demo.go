package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)


// 示例：扫描
func main() {
	// 逗号分隔的字符串，最后一项为空
	const input = "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 定义匹配函数（查找逗号分隔的字符串）
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if atEOF {
			// 告诉 Scanner 扫描结束。
			return 0, data, bufio.ErrFinalToken
		} else {
			// 告诉 Scanner 没找到匹配项，让 Scan 填充缓存后再次扫描。
			return 0, nil, nil
		}
	}
	// 指定匹配函数
	scanner.Split(onComma)
	// 开始扫描
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	// 检查是否因为遇到错误而结束
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
