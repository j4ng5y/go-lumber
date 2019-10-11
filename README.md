# Go-Lumber
!["Photo by Dominika Roseclay from Pexels"](https://images.pexels.com/photos/1239420/pexels-photo-1239420.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260)

!["Go Report Card"](https://goreportcard.com/badge/github.com/j4ng5y/go-lumber) ![Code Climate maintainability](https://img.shields.io/codeclimate/maintainability-percentage/j4ng5y/go-lumber) ![Code Climate coverage](https://img.shields.io/codeclimate/coverage/j4ng5y/go-lumber?label=Test%20Coverage)

Go lumber is a simple, yet opinionated little logging library that I seem to write into a lot of my packages, so I decided to make it a library instead of re-writing it every time.

Documentation is available via godoc here: [https://godoc.org/github.com/j4ng5y/go-lumber](https://godoc.org/github.com/j4ng5y/go-lumber)

## Installation

`go get -u github.com/j4ng5y/go-lumber`

## Usage

To use this library, you can do something like the following:

```go
package main

import "github.com/j4ng5y/go-lumber"

func main() {
    log := lumber.New()

    data, err := ErroniousFunction()
    if err != nil {
        log.Errorf("this function failed: %v", err)
    }
}
```

You can also set a log file in addition to just writing to STDOUT by doing the following:

```go
...
log := lumber.New()

log.SetLogFile("debug", "debug.log")
...
```

and then use the library like normal.