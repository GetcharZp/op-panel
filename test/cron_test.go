package test

import (
	"github.com/robfig/cron/v3"
	"log"
	"testing"
)

func TestCron(t *testing.T) {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("*/2 * * * * *", func() {
		log.Println("RUN")
	})
	if err != nil {
		t.Fatal(err)
	}
	c.Start()
	defer c.Stop()
	select {}
}
