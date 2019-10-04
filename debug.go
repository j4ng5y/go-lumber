package lumber

// Debug takes a message string and logs that message
//
// Arguements:
//     msg (string): The message to log
//
// Returns:
//     None
func (L Lumber) Debug(msg string) {
	L.DebugLog.Print(msg)
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
}
