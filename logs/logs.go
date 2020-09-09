package main

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	//USE THIS if you want to write on specific file, not the os.Stderr
	//file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	file := os.Stderr

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	log.Println("Teste log")

	//Custom Logs
	InfoLogger.Println("Starting the application...")
	InfoLogger.Println("Something noteworthy happened")
	WarningLogger.Println("There is something you should know about")
	ErrorLogger.Println("Something went wrong")

	//There are a few of robust logging frameworks, like LOGRUS wich supports log levels and other features
}
