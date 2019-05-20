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

import (
	f "fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestNewFileLogger(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    FileLogger
		wantErr bool
	}{
		{
			"TestNewFileLogger #1-1",
			args{"./flog.log"},
			&logger{1, "./flog.log", nil},
			false,
		},
		{
			"TestNewFileLogger #1-2",
			args{""},
			&logger{1, "", nil},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFileLogger(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFileLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("NewFileLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogger_printToLog(t *testing.T) {
	log := &logger{_Info, "./flog.log", nil}
	tNow := time.Now()
	timestamp := f.Sprintf("[%02d.%02d.%d %02d:%02d:%02d]: %s ", tNow.Day(), tNow.Month(), tNow.Year(), tNow.Hour(), tNow.Minute(), tNow.Second(), "INFO")
	printArg := "This is an informational message"
	type args struct {
		formatString string
		args         []interface{}
	}
	tests := []struct {
		name    string
		l       *logger
		args    args
		want    int
		wantErr bool
	}{
		{
			"TestNewFileLogger #2",
			log,
			args{printArg, nil},
			len(timestamp) + len(printArg) + 1, // "+1", because of the appended "\n" newline
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.printToLog(tt.args.formatString, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Logger.printToLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Logger.printToLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_openHandle(t *testing.T) {
	fHandle, fHandleErr := os.Open("./flog.log")
	if fHandleErr != nil {
		return
	}
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *os.File
		wantErr bool
	}{
		{
			"TestNewFileLogger #3",
			args{"./flog.log"},
			fHandle,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := openHandle(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("openHandle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// use a weaker check here...
			if !(reflect.TypeOf(got).String() == reflect.TypeOf(tt.want).String()) {
				t.Errorf("openHandle() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_logger_PrintInfo(t *testing.T) {
	log := logger{_Info, "./flog.log", nil}
	tNow := time.Now()
	timestamp := f.Sprintf("[%02d.%02d.%d %02d:%02d:%02d]: %s ", tNow.Day(), tNow.Month(), tNow.Year(), tNow.Hour(), tNow.Minute(), tNow.Second(), "INFO")
	printArg := "This is an informational message"
	type args struct {
		formatString string
		args         []interface{}
	}
	tests := []struct {
		name    string
		l       logger
		args    args
		want    int
		wantErr bool
	}{
		{
			"TestNewFileLogger #4",
			log,
			args{printArg, nil},
			len(timestamp) + len(printArg) + 1, // "+1", because of the appended "\n" newline
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.PrintInfo(tt.args.formatString, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("logger.PrintInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("logger.PrintInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_PrintDebug(t *testing.T) {
	log := logger{_Debug, "./flog.log", nil}
	tNow := time.Now()
	timestamp := f.Sprintf("[%02d.%02d.%d %02d:%02d:%02d]: %s ", tNow.Day(), tNow.Month(), tNow.Year(), tNow.Hour(), tNow.Minute(), tNow.Second(), "DEBUG")
	printArg := "This is a debugging message"
	type args struct {
		formatString string
		args         []interface{}
	}
	tests := []struct {
		name    string
		l       logger
		args    args
		want    int
		wantErr bool
	}{
		{
			"TestNewFileLogger #5",
			log,
			args{printArg, nil},
			len(timestamp) + len(printArg) + 1, // "+1", because of the appended "\n" newline
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.PrintDebug(tt.args.formatString, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("logger.PrintDebug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("logger.PrintDebug() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_PrintWarning(t *testing.T) {
	log := logger{_Warning, "./flog.log", nil}
	tNow := time.Now()
	timestamp := f.Sprintf("[%02d.%02d.%d %02d:%02d:%02d]: %s ", tNow.Day(), tNow.Month(), tNow.Year(), tNow.Hour(), tNow.Minute(), tNow.Second(), "WARNING")
	printArg := "This is a warning message"
	type args struct {
		formatString string
		args         []interface{}
	}
	tests := []struct {
		name    string
		l       logger
		args    args
		want    int
		wantErr bool
	}{
		{
			"TestNewFileLogger #6",
			log,
			args{printArg, nil},
			len(timestamp) + len(printArg) + 1, // "+1", because of the appended "\n" newline
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.PrintWarning(tt.args.formatString, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("logger.PrintWarning() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("logger.PrintWarning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_PrintError(t *testing.T) {
	log := logger{_Error, "./flog.log", nil}
	tNow := time.Now()
	timestamp := f.Sprintf("[%02d.%02d.%d %02d:%02d:%02d]: %s ", tNow.Day(), tNow.Month(), tNow.Year(), tNow.Hour(), tNow.Minute(), tNow.Second(), "ERROR")
	printArg := "This is an error message"
	type args struct {
		formatString string
		args         []interface{}
	}
	tests := []struct {
		name    string
		l       logger
		args    args
		want    int
		wantErr bool
	}{
		{
			"TestNewFileLogger #7-1",
			log,
			args{printArg, nil},
			len(timestamp) + len(printArg) + 1, // "+1", because of the appended "\n" newline
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.PrintError(tt.args.formatString, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("logger.PrintError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("logger.PrintError() = %v, want %v", got, tt.want)
			}
		})
	}
}
