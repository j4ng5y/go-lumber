// lumber is a highly optinionated logging package for Go that I seem to use in all of my packages so, it was time to make it a module

package lumber

import (
	"log"
	"os"
)

// Lumber is a struct to hold loggers
type Lumber struct {
	DebugLog *log.Logger
	InfoLog  *log.Logger
	WarnLog  *log.Logger
	ErrorLog *log.Logger
	FatalLog *log.Logger
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
		WarnLog:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLog: log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
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
func (L Lumber) SetLogFile(filename string, l *log.Logger) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	l.SetOutput(f)
	return nil
}

// Debug takes a message and logs that message
func (L Lumber) Debug(msg string) {
	L.DebugLog.Println(msg)
}

// Debugf takes a formatted message and logs that message
func (L Lumber) Debugf(format string, v ...interface{}) {
	L.DebugLog.Printf(format, v...)
}

// Info takes a message and logs that message
func (L Lumber) Info(msg string) {
	L.InfoLog.Println(msg)
}

// Infof takes a formatted message and logs that message
func (L Lumber) Infof(format string, v ...interface{}) {
	L.InfoLog.Printf(format, v...)
}

// Warn takes a message and logs that message as a warning message
func (L Lumber) Warn(msg string) {
	L.WarnLog.Println(msg)
}

// Warnf takes a formatted message and logs that message
func (L Lumber) Warnf(format string, v ...interface{}) {
	L.WarnLog.Printf(format, v...)
}

// Error takes a message and logs that message as a erronious message
func (L Lumber) Error(msg string) {
	L.ErrorLog.Println(msg)
}

// Errorf takes a formatted message and logs that message
func (L Lumber) Errorf(format string, v ...interface{}) {
	L.ErrorLog.Printf(format, v...)
}

// Fatal tages a message and logs that message as a fatal message
func (L Lumber) Fatal(msg string) {
	L.FatalLog.Fatal(msg)
}

// Fatalf takes a formatted message and logs that message
func (L Lumber) Fatalf(format string, v ...interface{}) {
	L.FatalLog.Printf(format, v...)
}
