package controller

import "time"

func (watcher *WatchDog) Start() {
	for _, job := range watcher.Jobs {
		jobItem := job
		go func() {
			for {
				time.Sleep(time.Duration(jobItem.MonitoringPeriod) * time.Second)
				isRunning, err := ProcessIsRunning(jobItem.ProcessName)
				if err != nil {
					continue
				}

				if !isRunning {
					_, _ = ExecCmd(jobItem.ProcessStartupScript)
				}
			}
		}()
	}
}
