
# How To Use

基于 go 模块的方式，使用这个库
```
$ go mod init test
$ go get github.com/rootkiter/FLog
```

下面是一段测试代码

```
package main
import (
    "time"
    "github.com/rootkiter/FLog"
)

func main() {
    loghandle := FLog.FLog{}

    loghandle.Init(
        "/tmp",        // Log directory
        "logtest",     // Log name (file name prompt)
        1*1000         // Create a new file every second
    )
    loghandle.WriteLogString("HELLO") // Record a log

    time.Sleep(2*1000*1000*1000)      // Wait 2 seconds

    loghandle.WriteLogString("World") // Record the next log in a new file.

    loghandle.Close()
}
```
