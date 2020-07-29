package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	c.AddFunc("@every 1s", func() {
		fmt.Println("tick every 1 second")
	})

	c.Start()
	time.Sleep(time.Second * 5)
}
