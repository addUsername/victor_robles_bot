package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"regexp"
)

func handleError(err error, errType string) {
	if err != nil {
		if errType == "fatal" {
			log.Fatalln(err)
		} else if errType == "panic" {
			log.Panicln(err)
		} else if errType == "error" {
			log.Errorln(err)
		} else if errType == "warn" {
			log.Warnln(err)
		} else if errType == "debug" {
			log.Debugln(err)
		} else if errType == "info" {
			log.Infoln(err)
		}
	}
}

func handleError403(bot *tb.Bot, chatID tb.ChatID, err error) bool {
	const textTooLongErrMsg = "El texto que te quiero enviar es un poco largo, " +
		"ábreme un privado a @victor_robles_bot para poder mandártelo."
	if err != nil && is403(err) {
		sendMessage(bot, chatID, textTooLongErrMsg)
		handleError(err, "error")
		return true
	}
	return false
}

func is403(err error) bool {
	regexExpr := "403"
	regex, errCompile := regexp.Compile(regexExpr)
	handleError(errCompile, "error")
	// err.Error() is an interface for the String contents of the Error type.
	match := regex.MatchString(err.Error())
	return match
}
