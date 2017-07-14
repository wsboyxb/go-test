package main

import (
	"fmt"
	"time"

	"github.com/zemirco/uid"
)

func task(t time.Time) {
	fmt.Println("time is", t)
}

func main() {
	fmt.Println("begin")
	/*cron.NewDailyJob(cron.ANY, cron.ANY, 5, task)
	cron.NewDailyJob(cron.ANY, cron.ANY, 1, func(t time.Time) {
		fmt.Println(t)
	})
	time.Sleep(time.Second * 200)*/
	id := uid.New(10)
	fmt.Println("uid is: ", id)
}
