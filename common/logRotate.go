package common

import (
	"log"
	"os"
	"time"
)

func SetLogOut() *os.File {
	logFile, err := os.OpenFile(LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	return logFile
}
func SetLogOutTesting() *os.File {
	logFile, err := os.OpenFile(LogPathTesting, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	return logFile
}
func LogRotate() {
	fileInfo, err := os.Stat(LogPath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatal("Log file checking failed")
	}
	if fileInfo.Size() < MaxLogFileSize {
		return
	}
	err = os.Rename(LogPath, LogPath+"."+time.Now().Format(time.RFC3339))
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(LogPath)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
