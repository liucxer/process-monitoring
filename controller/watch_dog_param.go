package controller

type WatchDog struct {
	Jobs []WatchDogJob `json:"jobs"`
}

type WatchDogJob struct {
	ProcessName          string `json:"processName"`
	MonitoringPeriod     int64  `json:"monitoringPeriod"`
	ProcessStartupScript string `json:"processStartupScript"`
}
