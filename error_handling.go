package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"regexp"
)

func handleError(err error, fatal bool) {
	if err != nil {
		if fatal {
			log.Fatal(err)
		}
		log.Error(err)
	}
}

func handleError403(bot *tb.Bot, chatID tb.ChatID, err error) bool {
	if err != nil && is403(err) {
		sendMessage(bot, chatID, "El texto que te quiero enviar es un poco largo, no te lo puedo enviar si no me abres un privado @victor_robles_bot")
		handleError(err, false)
		return true
	}
	return false
}

func is403(err error) bool {
	regexExpr := "403"
	regex, errCompile := regexp.Compile(regexExpr)
	handleError(errCompile, false)
	match := regex.MatchString(err.Error())
	return match
}
