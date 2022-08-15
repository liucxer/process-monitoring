package controller

import (
	"github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
	"time"
)

func ExecCmd(name string) (string, error) {
	var (
		result string
	)
	startTime := time.Now()
	logrus.Debugf("ExecCmd start. name:%s", name)
	defer func() {
		cost := time.Now().Sub(startTime).Seconds()
		logrus.Debugf("ExecCmd end.   cmd:\"%s\", cost:%fs, result:\n%s", name, cost, result)
	}()

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", name)
		out, err := cmd.CombinedOutput()
		if err != nil {
			logrus.Errorf("exec name:%s error:%v", name, err)
			return "", err
		}
		logrus.Warningf("exec name:%s success", name)
		result = string(out)
		return string(out), nil
	} else {
		cmd := exec.Command("bash", "-c", name)
		out, err := cmd.CombinedOutput()
		if err != nil {
			logrus.Errorf("exec name:%s error:%v", name, err)
			return "", err
		}
		logrus.Warningf("exec name:%s success", name)
		result = string(out)
		return string(out), nil
	}
}
