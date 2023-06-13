package main

import (
	"github.com/degreane/ezekod.com/middleware/ezelogger"
)

// Init function runs before any system startup
func init() {
	ezelogger.SetLogger("ezekod.log")
	ezelogger.Ezelogger.Println("Starting System Init in ezelogger")
}

func main() {

}
