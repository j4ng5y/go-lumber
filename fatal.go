package lumber

import "fmt"

// Fatal takes a message string and logs that message as a fatal message
//
// Arguements:
//     msg (string): The message to log
//
// Returns:
//     None
func (L Lumber) Fatal(msg string) {
	L.FatalLog.Fatal(msg)

	// If file logging is set, also log to the file
	if L.FatalFileLogger != nil {
		L.FatalFileLogger.Println(msg)
	}
}

// Fatalln takes a message string and logs that message as a fatal message with a newline inluded
//
// Arguements:
//     msg (string): The message to log
//
// Returns:
//     None
func (L Lumber) Fatalln(msg string) {
	L.DebugLog.Fatalln(msg)

	// If file logging is set, also log to the file
	if L.FatalFileLogger != nil {
		L.FatalFileLogger.Println(msg)
	}
}

// Fatalf takes a formatted message string and logs that message as a fatal message
//
// Arguements:
//     format (string): The formatted message to log
//     v (...interface{}): any number of variables to use to format the message
//
// Returns:
//     None
func (L Lumber) Fatalf(format string, v ...interface{}) {
	L.FatalLog.Fatalf(format, v...)

	// If file logging is set, also log to the file
	if L.FatalFileLogger != nil {
		L.FatalFileLogger.Println(fmt.Sprintf(format, v...))
	}
}
