/*
Package flog provides the ability to do file logging (hence the name)

	File: flog.go
	Version: 2.0
	Synopsis: provides file logging for the other packages
	Author: Jeff Wagner
	E-mail: hackmac89@filmdatenbank-manager.de
	date: 05/13/2019
	last modification date: 05/14/2019
*/

//Package flog provides the ability to do file logging (hence the name)
package flog

//Importe
import (
	"errors"
	f "fmt"
	"os"
	"sync"
	"time"
)

var mutex sync.Mutex

type loggingLevel uint

const (
	_Info loggingLevel = 1 << iota
	_Debug
	_Warning
	_Error
)

//FileLogger ...
type FileLogger interface {
	PrintInfo(formatString string, args ...interface{}) (int, error)
	PrintDebug(formatString string, args ...interface{}) (int, error)
	PrintWarning(formatString string, args ...interface{}) (int, error)
	PrintError(formatString string, args ...interface{}) (int, error)
}

//logger ...
type logger struct {
	level   loggingLevel
	fPath   string
	fHandle *os.File
}

//NewFileLogger returns a pointer to our logger
func NewFileLogger(filePath string) (FileLogger, error) {
	if len(filePath) > 0 {
		return &logger{_Info, filePath, nil}, nil
	}
	return nil, errors.New("ERROR(flog.go -> NewFileLogger): No path was given")
}

//PrintToLog writes a given message and a timestamp into a logfile (I´m usually not checking for errors when calling this function, so errors are quite silent)
func (l *logger) printToLog(formatString string, args ...interface{}) (int, error) {
	// Declarations
	var bytesWritten int
	var handleErr error
	var level string

	mutex.Lock() // lock access
	defer mutex.Unlock()
	t := time.Now()
	switch l.level {
	case _Info:
		level = "INFO"
	case _Debug:
		level = "DEBUG"
	case _Warning:
		level = "WARNING"
	case _Error:
		level = "ERROR"
	default:
		level = "INFO"
	}
	timestamp := f.Sprintf("[%02d.%02d.%d %02d:%02d:%02d]: %s ", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second(), level)
	// Re-obtain file fHandle
	l.fHandle, handleErr = openHandle(l.fPath)
	if handleErr != nil {
		return 0, handleErr
	}

	defer l.fHandle.Close() // defer closing of the file handler
	// Write timestamp and message to file
	bytesWritten, err := l.fHandle.WriteString(f.Sprintf(timestamp+formatString+"\n", args...))
	if err != nil {
		return 0, err
	}
	return bytesWritten, nil
}

//openHandle returns a handle for the opened file or an error otherwise
func openHandle(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0744) // Use "Append" to append messages to an existing logfile or create a new one if needed
}

//PrintInfo prints a standard informational message
func (l logger) PrintInfo(formatString string, args ...interface{}) (int, error) {
	l.level = _Info
	return l.printToLog(formatString, args...)
}

//PrintDebug prints a debugging message
func (l logger) PrintDebug(formatString string, args ...interface{}) (int, error) {
	l.level = _Debug
	return l.printToLog(formatString, args...)
}

//PrintWarning prints a warning message
func (l logger) PrintWarning(formatString string, args ...interface{}) (int, error) {
	l.level = _Warning
	return l.printToLog(formatString, args...)
}

//PrintError prints an error message
func (l logger) PrintError(formatString string, args ...interface{}) (int, error) {
	l.level = _Error
	return l.printToLog(formatString, args...)
}
