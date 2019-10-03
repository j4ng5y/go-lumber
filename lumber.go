package lumber

import (
	"log"
	"os"
)

// Lumber is a struct to hold loggers
type Lumber struct {
	InfoLog  *log.Logger
	WarnLog  *log.Logger
	ErrorLog *log.Logger
}

// New returns a new instance of Lumber
func New() *Lumber {
	return &Lumber{
		InfoLog:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
		WarnLog:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLog: log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info takes a message and logs that message
func (L Lumber) Info(msg interface{}) {
	L.InfoLog.Println(msg.(string))
}

// Warn takes a message and logs that message as a warning message
func (L Lumber) Warn(msg interface{}) {
	L.WarnLog.Println(msg.(string))
}

// Error takes a message and logs that message as a erronious message
func (L Lumber) Error(msg interface{}) {
	L.ErrorLog.Println(msg.(string))
}

// Fatal tages a message and logs that message as a fatal message
func (L Lumber) Fatal(msg interface{}) {
	L.ErrorLog.Fatal(msg.(string))
}
