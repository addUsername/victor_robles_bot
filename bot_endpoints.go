package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
)

// This is where endpoints are declared
func loadEndpoints(bot *tb.Bot) {
	handleEndpoint(bot, "!help", loadTextFile("help"), false)
	handleEndpoint(bot, "!ayuda", loadTextFile("help"), false)

	handleEndpoint(bot, "!cursos java ash", loadTextFile("cursosjavaash"), false)
	handleEndpoint(bot, "!recomendaciones", loadTextFile("recomendaciones"), false)

	handleEndpoint(bot, "!acceso", loadTextFile("acceso"), false)
	handleEndpoint(bot, "!acceso links", loadTextFile("acceso_links"), false)
	handleEndpoint(bot, "!asignaturas primero", loadTextFile("asignaturas_primero"), true)
	handleEndpoint(bot, "!mensaje privado", "Prueba", true)
}

func loadTextFile(textFileName string) string {
	folder := "cursosinfo/"
	fileName := folder + textFileName + ".txt"
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Warnln(err)
	}
	return string(fileBytes)
}
