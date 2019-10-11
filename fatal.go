// Copyright 2019 Jordan Gregory (https://github.com/j4ng5y)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0

package lumber

import (
	"fmt"
	"os"
)

// Fatal takes a message string and logs that message as a fatal message
//
// Arguments:
//     msg (string):   The message to log
//     exitCode (int): The exit code to exit the application with
//
// Returns:
//     None
func (L Lumber) Fatal(msg string, exitCode int) {
	L.FatalLog.Print(msg)
	os.Exit(exitCode)

	// If file logging is set, also log to the file
	if L.FatalFileLogger != nil {
		L.FatalFileLogger.Println(msg)
	}
}

// Fatalln takes a message string and logs that message as a fatal message with a newline inluded
//
// Arguments:
//     msg (string):   The message to log
//     exitCode (int): The exit code to exit the application with
//
// Returns:
//     None
func (L Lumber) Fatalln(msg string, exitCode int) {
	L.FatalLog.Println(msg)
	os.Exit(exitCode)

	// If file logging is set, also log to the file
	if L.FatalFileLogger != nil {
		L.FatalFileLogger.Println(msg)
	}
}

// Fatalf takes a formatted message string and logs that message as a fatal message
//
// Arguments:
//     exitCode (int):     The exit code to exit the application with
//     format (string):    The formatted message to log
//     v (...interface{}): Any number of variables to use to format the message
//
// Returns:
//     None
func (L Lumber) Fatalf(exitCode int, format string, v ...interface{}) {
	L.FatalLog.Printf(format, v...)
	os.Exit(exitCode)

	// If file logging is set, also log to the file
	if L.FatalFileLogger != nil {
		L.FatalFileLogger.Println(fmt.Sprintf(format, v...))
	}
}
