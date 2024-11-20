package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/robfig/cron/v3"
)

func Notification(title string) {
    err := exec.Command("notify-send", "Todo Reminder", fmt.Sprintf("Reminder: '%s' is due in 30 minutes!", title)).Run()
    if err != nil {
        log.Println("Error sending notification:", err)
    }
}


func RunNotifications(todos *Todos) {
	log.Println("Notification system started...")

	now := time.Now()
	c := cron.New(cron.WithSeconds())


	fmt.Println("Notification system started...")


	for _, t := range *todos {
		if t.Deadline != nil {
			notifyTime := t.Deadline.Add(-45 * time.Minute)
			if notifyTime.After(now) {
				log.Printf("Scheduling notification for '%s' at %v\n", t.Title, notifyTime)
				_, err := c.AddFunc(fmt.Sprintf("%d %d %d %d %d *",
					notifyTime.Second(),
					notifyTime.Minute(),
					notifyTime.Hour(),
					notifyTime.Day(),
					int(notifyTime.Month())),
					func(title string) func() {
						return func() {
							Notification(title)
						}
					}(t.Title))
				if err != nil {
					log.Printf("Error scheduling notification for '%s': %v", t.Title, err)
				}
			}
		}
	}

	c.Start()
	log.Println("Notification system is running in the background.")
}



func ParseLocalTime(time_string time.Time) time.Time {
	localLocation := time.Now().Location()
	localDeadline := time.Date(
		time_string.Year(),
		time_string.Month(),
		time_string.Day(),
		time_string.Hour(),
		time_string.Minute(),
		time_string.Second(),
		time_string.Nanosecond(),
		localLocation,
	)

	return localDeadline

}