// Package 能做成套件給其他Go import的最小包
// go.mod 中 module 應該為你的github路徑，這樣可以避免引用名稱不同的問題
// 檔名不要使用main.go 其餘應該都可以
package GoHelloWorldForImport // 這個名稱將會是別人使用你的工具使用的名字 EX GoHelloWorldForImport.SayIt()

import "fmt"

func SayIt() { // 這個名稱將會是別人使用你工具的函數名 EX GoHelloWorldForImport.SayIt()
	fmt.Println("Hello world")
}
