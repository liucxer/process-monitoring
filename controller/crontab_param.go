package controller

type Crontab struct {
	Jobs []CrontabJob `json:"jobs"`
}

type CrontabJob struct {
	JobName   string `json:"jobName"`
	JobSpec   string `json:"jobSpec"`
	JobScript string `json:"jobScript"`
}
