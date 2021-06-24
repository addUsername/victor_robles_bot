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
	handleEndpoint(bot, "!cursos max", loadTextFile("cursosmax"), false)
	handleEndpoint(bot, "!recomendaciones", loadTextFile("recomendaciones"), false)

	handleEndpoint(bot, "!acceso", loadTextFile("acceso"), false)
	handleEndpoint(bot, "!acceso links", loadTextFile("acceso_links"), false)
	handleEndpoint(bot, "!asignaturas primero", loadTextFile("asignaturas_primero"), true)
	handleEndpoint(bot, "!asignaturas segundo daw", loadTextFile("asignaturas_segundo_daw"), true)
}

func loadTextFile(textFileName string) string {
	folder := "cursosinfo/"
	fileName := folder + textFileName + ".txt"
	fileBytes, err := ioutil.ReadFile(fileName)
	handleError(err, "warn")
	return string(fileBytes)
}
