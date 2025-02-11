package common

import (
	"log"
	"os"
	"time"
)

func SetLogOut() *os.File {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	logDir := homeDir + "/rupamic_arch_log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	logFile, err := os.OpenFile(logDir+"/"+LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
