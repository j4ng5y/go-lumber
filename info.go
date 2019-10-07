// Copyright 2019 Jordan Gregory (https://github.com/j4ng5y)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0

package lumber

import "fmt"

// Info takes a message string and logs that message as an info message
//
// Arguments:
//     msg (string): The message to log
//
// Returns:
//     None
func (L Lumber) Info(msg string) {
	L.InfoLog.Print(msg)

	// If file logging is set, also log to the file
	if L.InfoFileLogger != nil {
		L.InfoFileLogger.Println(msg)
	}
}

// Infoln takes a message string and logs that message with a newline inluded
//
// Arguments:
//     msg (string): The message to log
//
// Returns:
//     None
func (L Lumber) Infoln(msg string) {
	L.InfoLog.Println(msg)

	// If file logging is set, also log to the file
	if L.InfoFileLogger != nil {
		L.InfoFileLogger.Println(msg)
	}
}

// Infof takes a formatted message string and logs that message as an info message
//
// Arguments:
//     format (string): The formatted message to log
//     v (...interface{}): any number of variables to use to format the message
//
// Returns:
//     None
func (L Lumber) Infof(format string, v ...interface{}) {
	L.InfoLog.Printf(format, v...)

	// If file logging is set, also log to the file
	if L.InfoFileLogger != nil {
		L.InfoFileLogger.Println(fmt.Sprintf(format, v...))
	}
}
