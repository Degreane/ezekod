package main

import "github.com/degreane/ezekod.com/middleware/ezelogger"

// Init function runs before any system startup
func init() {
	ezelogger.SetLogger("ezekod.log")
	ezelogger.Ezelogger.Println("Starting System Init in ezelogger")
	err := DB.Connect()
	if err != nil {
		ezelogger.Ezelogger.Fatalf("%+v", err)
	}
	ezelogger.Ezelogger.Printf("%+v", DB)
	DB.CloseDB()
}

func main() {

}
