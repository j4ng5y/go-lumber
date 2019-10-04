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
func (L Lumber) SetLogFile(filename string, l *log.Logger) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	l.SetOutput(f)
	return nil
}
