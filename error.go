package lumber

// Error takes a message string and logs that message as a erronious message
//
// Arguements:
//     msg (string): The message to log
//
// Returns:
//     None
func (L Lumber) Error(msg string) {
	L.ErrorLog.Print(msg)
}

// Errorln takes a message string and logs that message with a newline inluded
//
// Arguements:
//     msg (string): The message to log
//
// Returns:
//     None
func (L Lumber) Errorln(msg string) {
	L.ErrorLog.Println(msg)
}

// Errorf takes a formatted message string and logs that message as an erronious message
//
// Arguements:
//     format (string): The formatted message to log
//     v (...interface{}): any number of variables to use to format the message
//
// Returns:
//     None
func (L Lumber) Errorf(format string, v ...interface{}) {
	L.ErrorLog.Printf(format, v...)
}
