// Package 能做成套件給其他Go import的最小包
package GoHelloWorldForImport2 // 這個名稱將會是別人使用你的工具使用的名字 EX GoHelloWorldForImport.SayIt()

import "fmt"

func SayIt() { // 這個名稱將會是別人使用你工具的函數名 EX GoHelloWorldForImport.SayIt()
	fmt.Println("Hello world")
}
