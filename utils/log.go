package utils

import "log"

var (
	print = false
	lll   = true
)

func Printf(format string, v ...interface{}) {
	if print {
		log.Printf(format, v...)

	}
}

func Logf(format string, v ...interface{}) {
	if lll {
		log.Printf(format, v...)
	}
}
