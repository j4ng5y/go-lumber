package lumber

import (
	"bytes"
	"log"
	"regexp"
	"testing"

	"github.com/matryer/is"
)

func Test_Error(t *testing.T) {
	is := is.New(t)
	var b bytes.Buffer
	T := &Lumber{
		ErrorLog: log.New(&b, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	T.Error("test")

	line := b.String()
	line = line[0 : len(line)-1]
	pattern := "^\\[ERROR\\].+?test$"
	r, err := regexp.MatchString(pattern, line)
	is.NoErr(err)
	is.True(r)
}

func Test_Errorln(t *testing.T) {
	is := is.New(t)
	var b bytes.Buffer
	T := &Lumber{
		ErrorLog: log.New(&b, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	T.Errorln("test")

	line := b.String()
	line = line[0 : len(line)-1]
	pattern := "^\\[ERROR\\].+?test$"
	r, err := regexp.MatchString(pattern, line)
	is.NoErr(err)
	is.True(r)
}

func Test_Errorf(t *testing.T) {
	is := is.New(t)
	var b bytes.Buffer
	T := &Lumber{
		ErrorLog: log.New(&b, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	T.Errorf("%s", "test")

	line := b.String()
	line = line[0 : len(line)-1]
	pattern := "^\\[ERROR\\].+?test$"
	r, err := regexp.MatchString(pattern, line)
	is.NoErr(err)
	is.True(r)
}
