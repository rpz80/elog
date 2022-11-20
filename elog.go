package elog

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	severityDebug int = iota
	severityInfo
	severityWarning
	severityError
	severityCritical
)

var (
	globalSeverity int        = severityInfo
	mutex          sync.Mutex = sync.Mutex{}
	levelString               = flag.String("log-level", "-log-level={INFO", "DEBUG, WARNING, ERROR, CRITICAL}")
	logLevelSet    bool       = false
)

func setDefaultLevelUnsafe() {
	if !logLevelSet {
		switch *levelString {
		case "INFO":
			globalSeverity = severityInfo
		case "DEBUG":
			globalSeverity = severityDebug
		case "WARNING":
			globalSeverity = severityWarning
		case "ERROR":
			globalSeverity = severityError
		case "CRITICAL":
			globalSeverity = severityCritical
		}
		logLevelSet = true
	}
}

func severityToString(severity int) string {
	switch severity {
	case severityInfo:
		return "INFO    "
	case severityDebug:
		return "DEBUG   "
	case severityWarning:
		return "WARNING "
	case severityError:
		return "ERROR   "
	case severityCritical:
		return "CRITICAL"
	default:
		panic("Unknown severity")
	}
}

func createPrefix(severity int) string {
	return fmt.Sprintf(
		"%v %v", time.Now().Format("2006-01-02T15:04:05.0000"), severityToString((severity)))
}

func joinPrefixWithRest(prefix string, v ...any) []any {
	result := make([]any, 0)
	result = append(result, prefix)
	for _, s := range v {
		result = append(result, s)
	}
	return result
}

func Debug(format string, v ...any) {
	doLog(format, severityDebug, v...)
}

func Info(format string, v ...any) {
	doLog(format, severityInfo, v...)
}

func Warning(format string, v ...any) {
	doLog(format, severityWarning, v...)
}

func Error(format string, v ...any) {
	doLog(format, severityError, v...)
}

func Critical(format string, v ...any) {
	mutex.Lock()
	defer mutex.Unlock()
	userMessage := fmt.Sprintf(format, v...)
	fmt.Fprintln(os.Stderr, createPrefix(severityCritical), userMessage)
	panic(userMessage)
}

func doLog(format string, severity int, v ...any) {
	mutex.Lock()
	defer mutex.Unlock()
	setDefaultLevelUnsafe()
	if globalSeverity <= severity {
		fmt.Fprintln(os.Stderr, createPrefix(severity), fmt.Sprintf(format, v...))
	}
}
