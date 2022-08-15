package controller

import (
	"fmt"
	"github.com/robfig/cron"
)

func (c *Crontab) Start() {
	for _, job := range c.Jobs {
		jobItem := job
		go func() {
			cronJob := cron.New()
			err := cronJob.AddFunc(jobItem.JobSpec, func() { ExecCmd(jobItem.JobScript) })
			if err != nil {
				panic(fmt.Sprintf("cronJob.AddFunc err:%v", err))
			}
			cronJob.Start()
		}()
	}
}
