# GoHellowWorldForImport
// Package 能做成套件給其他Go import的最小包

## 結論&要點
* go.mod 中 module 應該為你的github路徑，這樣可以避免引用名稱不同的問題
	* 小版本更新使用`tag`來上傳，若有破壞性更新就在module 最後方加上/v2, v3等字樣
* 檔名不要使用main.go 其餘應該都可以
* package會是別人取用工具包的名稱


## 在你的包中使用這個方式來引入
``` go
import (
	"github.com/limiu82214/GoHellowWorldForImport"
)
```

## 在你的包中這樣使用此套件
```go
GoHelloWorldForImport.SayIt()
```
