# cron
- https://darjun.github.io/2020/06/25/godailylib/cron/

mkdir cron && cd cron
go mod init github.com/darjun/go-daily-lib/cron

<!-- 安装cron，目前最新稳定版本为 v3： -->

go get -u github.com/robfig/cron/v3

```go

package main

import (
  "fmt"
  "time"

  "github.com/robfig/cron/v3"
)

func main() {
  c := cron.New()

  c.AddFunc("@every 1s", func() {
    fmt.Println("tick every 1 second")
  })

  c.Start()
  time.Sleep(time.Second * 5)
}
```
