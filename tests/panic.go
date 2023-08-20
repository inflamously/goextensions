package tests

import "log"

func StopPanic(message string) {
	if msg := recover(); msg == nil {
		log.Panic(message)
	} else {
		log.Print(msg)
	}
}
