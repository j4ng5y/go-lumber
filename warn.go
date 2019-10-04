package lumber

import "fmt"

// Warn takes a message string and logs that message as a warning message
//
// Arguements:
//     msg (string): The message to log
//
// Returns:
//     None
func (L Lumber) Warn(msg string) {
	L.WarnLog.Print(msg)

	// If file logging is set, also log to the file
	if L.WarnFileLogger != nil {
		L.WarnFileLogger.Println(msg)
	}
}

// Warnln takes a message string and logs that message with a newline inluded
//
// Arguements:
//     msg (string): The message to log
//
// Returns:
//     None
func (L Lumber) Warnln(msg string) {
	L.WarnLog.Println(msg)

	// If file logging is set, also log to the file
	if L.WarnFileLogger != nil {
		L.WarnFileLogger.Println(msg)
	}
}

// Warnf takes a formatted message string and logs that message as a warning message
//
// Arguements:
//     format (string): The formatted message to log
//     v (...interface{}): any number of variables to use to format the message
//
// Returns:
//     None
func (L Lumber) Warnf(format string, v ...interface{}) {
	L.WarnLog.Printf(format, v...)

	// If file logging is set, also log to the file
	if L.WarnFileLogger != nil {
		L.WarnFileLogger.Println(fmt.Sprintf(format, v...))
	}
}
