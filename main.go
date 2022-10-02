package main

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	file, _ := os.OpenFile(fmt.Sprintf("log/app_%d.log", time.Now().Minute()), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//v, _ := json.Marshal(a)

	logrus.Info("ilham")
	logrus.Info(time.Now())
	logrus.Error(errors.New("error test"))
}
