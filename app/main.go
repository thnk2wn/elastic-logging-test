package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/op/go-logging"
)

const (
	maxSleepMs = 30000
	minSleepMs = 5000
)

var (
	log = logging.MustGetLogger("main")
)

func main() {
	configureLogging()

	for {
		log.Infof("The time is now %v", time.Now().Format(time.RFC1123))
		sleepMs := rand.Int31n(maxSleepMs - minSleepMs + 1)
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	}
}

func configureLogging() {
	logging.SetFormatter(defaultLogFormat())

	writer, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}

	fileBackend := logging.AddModuleLevel(logging.NewLogBackend(writer, "localhost", 0))
	fileBackend.SetLevel(logging.DEBUG, "")

	logging.SetBackend(
		logging.NewBackendFormatter(fileBackend, defaultLogFileFormat()),
		logging.NewLogBackend(os.Stdout, "", 0))
}

func defaultLogFileFormat() logging.Formatter {
	return logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`)
}

func defaultLogFormat() logging.Formatter {
	return logging.MustStringFormatter(`%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x} %{message}`)
}
