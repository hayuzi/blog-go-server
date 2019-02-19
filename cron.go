package main

import (
	"github.com/robfig/cron"
	"log"
	"time"
)

func main_cron() {
	log.Println("Starting...")

	c := cron.New()
	// 添加模拟的定时任务
	c.AddFunc("*/5 * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
	})
	c.AddFunc("*/2 * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
	})

	c.Start()

	// 阻塞保证进程运行
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
