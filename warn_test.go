package lumber

import (
	"bytes"
	"log"
	"regexp"
	"testing"

	"github.com/matryer/is"
)

func Test_Warn(t *testing.T) {
	is := is.New(t)
	var b bytes.Buffer
	T := &Lumber{
		WarnLog: log.New(&b, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	T.Warn("test")

	line := b.String()
	line = line[0 : len(line)-1]
	pattern := "^\\[WARN\\].+?test$"
	r, err := regexp.MatchString(pattern, line)
	is.NoErr(err)
	is.True(r)
}

func Test_Warnln(t *testing.T) {
	is := is.New(t)
	var b bytes.Buffer
	T := &Lumber{
		WarnLog: log.New(&b, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	T.Warnln("test")

	line := b.String()
	line = line[0 : len(line)-1]
	pattern := "^\\[WARN\\].+?test$"
	r, err := regexp.MatchString(pattern, line)
	is.NoErr(err)
	is.True(r)
}

func Test_Warnf(t *testing.T) {
	is := is.New(t)
	var b bytes.Buffer
	T := &Lumber{
		WarnLog: log.New(&b, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	T.Warnf("%s", "test")

	line := b.String()
	line = line[0 : len(line)-1]
	pattern := "^\\[WARN\\].+?test$"
	r, err := regexp.MatchString(pattern, line)
	is.NoErr(err)
	is.True(r)
}
