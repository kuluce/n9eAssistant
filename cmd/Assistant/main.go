package main

import "n9eHelper/internal/logger"

func main() {
	logger.Infof("Hello, World!")
	logger.Errorf("Error message!")

	logger.Warn("warn message")
	logger.Infof("end")
}
