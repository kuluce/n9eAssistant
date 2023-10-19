package logger

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/toolkits/file"
)

var (
	// InfoLogger  *log.Logger
	// ErrorLogger *log.Logger
	// InfoLogger  *slog.Logger
	// ErrorLogger *slog.Logger
	InfoLogger  *glog.Logger
	ErrorLogger *glog.Logger
	ctx         = context.Background()
)

const (
	LogHOme     = "logs"
	LogFileName = "n9eHelper.log"
	InfoLog     = "info.log"
	ErrorLog    = "error.log"
)

func init() {
	logHome := os.Getenv("LOG_HOME")
	if logHome == "" {
		logHome = filepath.Join(file.SelfDir(), LogHOme)
	}

	// ensure log home exists
	if !file.IsExist(logHome) {
		err := os.MkdirAll(logHome, 0755)
		if err != nil {
			log.Fatal("Error creating log home: ", err)
			panic(err)
		}
	}

	// InfoLogger = glog.New()
	// ErrorLogger = glog.New()

	// InfoLogger.SetConfigWithMap(map[string]interface{}{
	// 	"stdoutPrint": true,
	// 	"level":       "all",
	// 	"path":        logHome,
	// 	"file":        InfoLog,
	// })
	// ErrorLogger.SetConfigWithMap(map[string]interface{}{
	// 	"stdoutPrint": true,
	// 	"level":       "all",
	// 	"path":        logHome,
	// 	"file":        ErrorLog,
	// })
	glog.SetConfigWithMap(map[string]interface{}{
		"stdoutPrint": true,
		"level":       "all",
		"path":        logHome,
		"file":        LogFileName,
	})

}

func Info(message string) {
	glog.Info(ctx, message)
}

func Infof(format string, args ...any) {
	glog.Infof(ctx, format, args...)
}

func Error(message string) {
	glog.Error(ctx, message)
}

func Errorf(format string, args ...any) {
	glog.Errorf(ctx, format, args...)
}

func Warn(message string) {
	glog.Warning(ctx, message)
}

func Warnf(format string, args ...any) {
	glog.Warningf(ctx, format, args...)
}

func Fatal(message string) {
	glog.Fatal(ctx, message)
}

func Fatalf(format string, args ...any) {
	glog.Fatalf(ctx, format, args...)
}

// func init() {
// 	logHome := os.Getenv("LOG_HOME")
// 	if logHome == "" {
// 		logHome = filepath.Join(file.SelfDir(), LogHOme)
// 	}

// 	// ensure log home exists
// 	if !file.IsExist(logHome) {
// 		err := os.MkdirAll(logHome, 0755)
// 		if err != nil {
// 			log.Fatal("Error creating log home: ", err)
// 			panic(err)
// 		}
// 	}

// 	infoLog := filepath.Join(logHome, InfoLog)
// 	errorLog := filepath.Join(logHome, ErrorLog)

// 	infoFile, err := os.OpenFile(infoLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatal("Error creating info log file: ", err)
// 	}
// 	errorFile, err := os.OpenFile(errorLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatal("Error creating error log file: ", err)
// 	}

// 	InfoLogger = log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
// 	ErrorLogger = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

// }

// func Info(message string) {
// 	InfoLogger.Println(message)
// }

// func Error(message string) {
// 	ErrorLogger.Println(message)
// }
