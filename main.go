package main

// Init function runs before any system startup

func main() {

	defer func() {
		// Close DB CONNECTION before exiting the system
		DB.CloseDB()
	}()
	// Initialize fibers app

}
