package lumber

import (
	"os"
	"strings"
	"testing"

	"github.com/matryer/is"
)

var T = New()

func Test_New(t *testing.T) {
	is := is.New(t)

	is.True(T.DebugLog != nil)
	is.True(T.InfoLog != nil)
	is.True(T.WarnLog != nil)
	is.True(T.ErrorLog != nil)
	is.True(T.FatalLog != nil)
	is.True(T.DebugFileLogger == nil)
	is.True(T.InfoFileLogger == nil)
	is.True(T.WarnFileLogger == nil)
	is.True(T.ErrorFileLogger == nil)
	is.True(T.FatalFileLogger == nil)
}

func Test_SetLogFile(t *testing.T) {
	is := is.New(t)

	files := map[string]string{
		"debug": "./debug.test.log",
		"info":  "./info.test.log",
		"warn":  "./warn.test.log",
		"error": "./error.test.log",
		"fatal": "./fatal.test.log",
	}

	for k, v := range files {
		err := T.SetLogFile(k, v)
		defer cleanup(v, t)

		is.NoErr(err)
		is.True(checkFileExists(v, t))
	}

	err := T.SetLogFile("this_does_not_exist", "./this_file_should_not_exist.log")
	is.True(strings.Contains(err.Error(), "this_does_not_exist is not a valid log level"))
	is.True(!checkFileExists("./this_file_should_not_exist.log", t))
}

func checkFileExists(filename string, t *testing.T) bool {
	s, err := os.Stat(filename)
	if err != nil {
		t.Logf("the file, %s, does not exist due to error: %v", filename, err)
		return false
	}
	if s.IsDir() {
		t.Logf("%s is a directory, not a file", filename)
		return false
	}
	return true
}

func cleanup(filename string, t *testing.T) {
	if err := os.Remove(filename); err != nil {
		t.Logf("Unable to clean up the file: %s", filename)
	}
}
