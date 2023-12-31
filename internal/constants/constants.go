package constants

import "os"

var SaveFile string

func init() {
	// Get home Directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	// Make sure ~/.config/SimpleOTP exists
	if _, err := os.Stat(homeDir + "/.config/SimpleOTP"); os.IsNotExist(err) {
		err = os.MkdirAll(homeDir+"/.config/SimpleOTP", 0700)
		if err != nil {
			panic(err)
		}
	}
	// Set the save file
	SaveFile = homeDir + "/.config/SimpleOTP/store.gob"
}
