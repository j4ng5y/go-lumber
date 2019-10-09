package lumber

import (
	"bytes"
	"log"
	"regexp"
	"testing"

	"github.com/matryer/is"
)

func Test_Debug(t *testing.T) {
	is := is.New(t)
	var b bytes.Buffer
	T := &Lumber{
		DebugLog: log.New(&b, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	T.Debug("test")

	line := b.String()
	line = line[0 : len(line)-1]
	pattern := "^\\[DEBUG\\].+?test$"
	r, err := regexp.MatchString(pattern, line)
	is.NoErr(err)
	is.True(r)
}

func Test_Debugln(t *testing.T) {
	is := is.New(t)
	var b bytes.Buffer
	T := &Lumber{
		DebugLog: log.New(&b, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	T.Debugln("test")

	line := b.String()
	line = line[0 : len(line)-1]
	pattern := "^\\[DEBUG\\].+?test$"
	r, err := regexp.MatchString(pattern, line)
	is.NoErr(err)
	is.True(r)
}

func Test_Debugf(t *testing.T) {
	is := is.New(t)
	var b bytes.Buffer
	T := &Lumber{
		DebugLog: log.New(&b, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	T.Debugf("%s", "test")

	line := b.String()
	line = line[0 : len(line)-1]
	pattern := "^\\[DEBUG\\].+?test$"
	r, err := regexp.MatchString(pattern, line)
	is.NoErr(err)
	is.True(r)
}
