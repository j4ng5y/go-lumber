// Copyright 2019 Jordan Gregory (https://github.com/j4ng5y)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0

package lumber

import "fmt"

// Error takes a message string and logs that message as a erronious message
//
// Arguements:
//     msg (string): The message to log
//
// Returns:
//     None
func (L Lumber) Error(msg string) {
	L.ErrorLog.Print(msg)

	// If file logging is set, also log to the file
	if L.ErrorFileLogger != nil {
		L.ErrorFileLogger.Println(msg)
	}
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

	// If file logging is set, also log to the file
	if L.ErrorFileLogger != nil {
		L.ErrorFileLogger.Println(msg)
	}
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

	// If file logging is set, also log to the file
	if L.ErrorFileLogger != nil {
		L.ErrorFileLogger.Println(fmt.Sprintf(format, v...))
	}
}
