package main

import (
	"process-monitoring/controller"
	"time"
)

func main() {
	go func() {
		watchDog := controller.WatchDog{
			Jobs: []controller.WatchDogJob{
				{
					ProcessName:          "processName",
					MonitoringPeriod:     1,
					ProcessStartupScript: "cd /root && nohup ./awesomeProject2 &",
				},
			},
		}
		watchDog.Start()
	}()

	go func() {
		crontab := controller.Crontab{
			Jobs: []controller.CrontabJob{
				{
					JobName:   "restart awesomeProject",
					JobSpec:   "@every 1m",
					JobScript: "killall awesomeProject2",
				},
			},
		}
		crontab.Start()
	}()

	for {
		time.Sleep(time.Second)
	}
}
