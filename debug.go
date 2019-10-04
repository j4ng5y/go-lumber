package lumber

import "fmt"

// Debug takes a message string and logs that message
//
// Arguements:
//     msg (string): The message to log
//
// Returns:
//     None
func (L Lumber) Debug(msg string) {
	L.DebugLog.Print(msg)

	// If file logging is set, also log to the file
	if L.DebugFileLogger != nil {
		L.DebugFileLogger.Println(msg)
	}
}

// Debugln takes a message string and logs that message with a newline inluded
//
// Arguements:
//     msg (string): The message to log
//
// Returns:
//     None
func (L Lumber) Debugln(msg string) {
	L.DebugLog.Println(msg)

	// If file logging is set, also log to the file
	if L.DebugFileLogger != nil {
		L.DebugFileLogger.Println(msg)
	}
}

// Debugf takes a formatted message string and logs that message as a debug message
//
// Arguements:
//     format (string): The formatted message to log
//     v (...interface{}): any number of variables to use to format the message
//
// Returns:
//     None
func (L Lumber) Debugf(format string, v ...interface{}) {
	L.DebugLog.Printf(format, v...)

	// If file logging is set, also log to the file
	if L.DebugFileLogger != nil {
		L.DebugFileLogger.Println(fmt.Sprintf(format, v...))
	}
}
