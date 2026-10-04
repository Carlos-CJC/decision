package main

import "fmt"

// render 按默认模板 {name}:{result} 生成单行输出。
func render(name, result string) string {
	return fmt.Sprintf("%s:%s", name, result)
}
