package additioner

import (
	"log"
	"os"
)

func Logg(Log bool) {
	if Log {
		file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			log.Println(err)
		}
		log.SetOutput(file)
	}
}

func Additioner(firstValue, secondValue float64) float64 {
	result := firstValue + secondValue
	log.SetPrefix("LOG: ")
	log.Println("Answer= ", result)
	return result
}
