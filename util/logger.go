package util

import (
	"log"
	"os"
)

// var logger *log.Logger
var logFile *os.File // 全局日志文件

// CreateLogger 创建一个自定义的日志示例，将日志输出到指定的文件中
func CreateLogger(logFileName string) *log.Logger {
	// 创建或打开日志文件
	var err error
	logFile, err = os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("无法打开日志文件: %v\n", err)
		return nil
	}

	// 创建自定义日志示例
	logger := log.New(logFile, "", log.LstdFlags)

	return logger
}

// CloseLogFile 关闭日志文件
func CloseLogFile() {
	if logFile != nil {
		logFile.Close()
	}
}
