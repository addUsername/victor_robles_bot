package main

import (
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)

var log = logrus.New()

func logSetup() {
	// This format is not to be changed to your current time, this is according to the constants at: https://golang.org/pkg/time/#pkg-constants
	log.SetFormatter(&logrus.TextFormatter{TimestampFormat: "02-01-2006 15:04:05 MST", FullTimestamp: true})
}

func logEndpointUsage(src *tb.Message, route string) {
	logString := src.Sender.Username + "(" + strconv.Itoa(src.Sender.ID) + "): Used -> " + route + " <-."
	log.Infoln(logString)
}
