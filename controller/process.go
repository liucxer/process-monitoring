package controller

import (
	"github.com/shirou/gopsutil/process"
	"github.com/sirupsen/logrus"
	"strings"
)

func ProcessIsRunning(processName string) (bool, error) {
	pids, err := process.Pids()
	if err != nil {
		logrus.Errorf("process.Pids err:%v", err)
		return false, err
	}
	for _, pid := range pids {
		process, err := process.NewProcess(pid)
		if err != nil {
			logrus.Errorf("process.NewProcess err:%v", err)
			return false, err
		}

		processNameItem, err := process.Name()
		if err != nil {
			logrus.Errorf("process.Name err:%v", err)
			return false, err
		}
		if strings.Contains(processNameItem, processName) {
			logrus.Infof("find process processNameItem:%s, processName:%s", processNameItem, processName)
			return true, nil
		}
	}
	return false, nil
}
