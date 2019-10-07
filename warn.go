// Copyright 2019 Jordan Gregory (https://github.com/j4ng5y)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0

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
