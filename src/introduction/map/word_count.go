package main

// fixme
import "golang.org/x/tour/wc"

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

// 字符串 s 中每个“单词”的个数
func main() {
	wc.Test(WordCount)
}
