package main

func init() {
	logSetup()
}

func main() {
	bot := createBot()
	loadEndpoints(bot)
	bot.Start()
}
