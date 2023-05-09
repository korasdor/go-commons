package logger

import "github.com/sirupsen/logrus"

func Print(msg ...interface{}) {
	logrus.Print(msg...)

	go writeToFile()
}

func Printf(format string, args ...interface{}) {
	logrus.Printf(format, args...)

	go writeToFile()
}

func Info(msg ...interface{}) {
	logrus.Info(msg...)

	go writeToFile()
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)

	go writeToFile()
}

func Warn(msg ...interface{}) {
	logrus.Warn(msg...)

	go writeToFile()
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)

	go writeToFile()
}

func Error(msg ...interface{}) {
	logrus.Error(msg...)

	go writeToFile()
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)

	go writeToFile()
}

func Fatal(msg ...interface{}) {
	logrus.Fatal(msg...)

	go writeToFile()
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)

	go writeToFile()
}

func writeToFile() {
	// TODO write to file feature
}
