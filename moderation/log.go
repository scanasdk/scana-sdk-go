package moderation

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

var logConf logConfType

type logConfType struct {
	level Level
}

func getDefaultLogConf() logConfType {
	return logConfType{
		level: LEVEL_INFO,
	}
}

func init() {
	logConf = getDefaultLogConf()
}

type Level int

const (
	LEVEL_ERROR Level = 400
	LEVEL_WARN  Level = 300
	LEVEL_INFO  Level = 200
	LEVEL_DEBUG Level = 100
)

var logLevelMap = map[Level]string{
	LEVEL_ERROR: "[ERROR]: ",
	LEVEL_WARN:  "[WARN]: ",
	LEVEL_INFO:  "[INFO]: ",
	LEVEL_DEBUG: "[DEBUG]: ",
}

func doLog(level Level, format string, v ...interface{}) {
	if logConf.level <= level {
		msg := fmt.Sprintf(format, v...)
		if _, file, line, ok := runtime.Caller(1); ok {
			index := strings.LastIndex(file, "/")
			if index >= 0 {
				file = file[index+1:]
			}
			msg = fmt.Sprintf("%s:%d|%s", file, line, msg)
		}
		prefix := logLevelMap[level]
		log.Printf("%s%s", prefix, msg)
	}
}

func checkAndLogErr(err error, level Level, format string, v ...interface{}) {
	if err != nil {
		doLog(level, format, v...)
	}
}
