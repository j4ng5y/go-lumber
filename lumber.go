// lumber is a highly optinionated logging package for Go that I seem to use in all of my packages so, it was time to make it a module

package lumber

import (
	"fmt"
	"log"
	"os"
)

// Lumber is a struct to hold loggers
type Lumber struct {
	DebugLog        *log.Logger
	DebugFileLogger *log.Logger
	InfoLog         *log.Logger
	InfoFileLogger  *log.Logger
	WarnLog         *log.Logger
	WarnFileLogger  *log.Logger
	ErrorLog        *log.Logger
	ErrorFileLogger *log.Logger
	FatalLog        *log.Logger
	FatalFileLogger *log.Logger
}

// New returns a new instance of Lumber
// This default behaviour will log everything to STDOUT
//
// Arguemnts:
//    None
//
// Returns:
//     (*Lumber): A pointer to the newly created instance of Lumber
func New() *Lumber {
	return &Lumber{
		DebugLog: log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLog:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
		WarnLog:  log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLog: log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		FatalLog: log.New(os.Stdout, "[FATAL] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// SetLogFile will set a log file for the provided logger
//
// Arguements:
//     filename (string): the name of the file to write
//     l (*log.Logger): the lumber logger to set this output file for
//
// Returns:
//    (error): an error if one exists, nil otherwise
func (L Lumber) SetLogFile(logLevel string, filename string) error {
	validLogLevels := []string{
		"debug",
		"info",
		"warn",
		"error",
		"fatal",
	}
	switch logLevel {
	case validLogLevels[0]:
		f, err := getFile(filename)
		if err != nil {
			return err
		}
		L.DebugFileLogger = log.New(f, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
		return nil
	case validLogLevels[1]:
		f, err := getFile(filename)
		if err != nil {
			return err
		}
		L.InfoFileLogger = log.New(f, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
		return nil
	case validLogLevels[2]:
		f, err := getFile(filename)
		if err != nil {
			return err
		}
		L.WarnFileLogger = log.New(f, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
		return nil
	case validLogLevels[3]:
		f, err := getFile(filename)
		if err != nil {
			return err
		}
		L.ErrorFileLogger = log.New(f, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
		return nil
	case validLogLevels[4]:
		f, err := getFile(filename)
		if err != nil {
			return err
		}
		L.FatalFileLogger = log.New(f, "[FATAL] ", log.Ldate|log.Ltime|log.Lshortfile)
		return nil
	default:
		return fmt.Errorf("%s is not a valid log level, must be one of: %v", logLevel, validLogLevels)
	}
}

func getFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}
