package ezelogger

import (
	"log"
	"os"
)

func SetLogger(fName string) {
	fD, err := os.OpenFile(fName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Could not Initiate Logger on File (%s) Err: %+v\n", fName, err)
	}
	Ezelogger = log.New(fD, "EzeKod ", log.LstdFlags)
}
