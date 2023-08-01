package utils

import "github.com/sirupsen/logrus"

func InitializeLogger() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.ForceColors = true
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
	logrus.SetReportCaller(true)
}
