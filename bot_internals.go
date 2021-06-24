package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

func createBot() *tb.Bot {
	// Creates a bot using the default URL setting (https://api.telegram.org), this can be overwritten if desired.
	bot, err := tb.NewBot(tb.Settings{
		Token: getToken(),
		// The 10 second timeout is the default and recommended setting, a lower number seems to affect performance.
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	handleError(err, "fatal")

	// Logging
	log.Infoln("Bot connected to Telegram Servers")

	// Change the default URL if you override the URL on tb.Settings.
	sendMessageToAdmin(bot, "Connection successful to "+tb.DefaultApiURL)

	return bot
}

func handleEndpoint(bot *tb.Bot, route string, message string, privateMsg bool) {
	bot.Handle(route, func(src *tb.Message) {
		// chatID represents the chat where the message was sent, which is not the same as the user than sent the message.
		chatID := tb.ChatID(src.Chat.ID)
		senderID := tb.ChatID(src.Sender.ID)

		if privateMsg {
			// It cannot be substituted for sendMessage() because it needs to retrieve the error in order to
			// check if the bot doesn't have permissions to send a message to the user (error 403).
			_, errSend := bot.Send(senderID, message, "html")
			if !handleError403(bot, chatID, errSend) {
				handleError(errSend, "error")
				sendMessage(bot, chatID, "Te lo he enviado por privado shur")
			}
			logEndpointUsage(src, route)
		} else {
			// This is the default way on handling endpoints.
			sendMessage(bot, chatID, message)
			logEndpointUsage(src, route)
		}
	})
}

func sendMessage(bot *tb.Bot, chatID tb.ChatID, message string) {
	// options: markdown doesn't seem to work, instead it uses html.
	_, err := bot.Send(chatID, message, "html")
	handleError(err, "error")
}

func sendMessageToAdmin(bot *tb.Bot, message string) {
	// This is hardcoded for now, since I'm the only admin.
	chatID := tb.ChatID(1099020633)
	sendMessage(bot, chatID, message)
	log.Warnln("Message: " + message + ". Sent to admins.")
}
