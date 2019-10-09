package lumber

import (
	"bytes"
	"log"
	"regexp"
	"testing"

	"github.com/matryer/is"
)

func Test_Info(t *testing.T) {
	is := is.New(t)
	var b bytes.Buffer
	T := &Lumber{
		InfoLog: log.New(&b, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	T.Info("test")

	line := b.String()
	line = line[0 : len(line)-1]
	pattern := "^\\[INFO\\].+?test$"
	r, err := regexp.MatchString(pattern, line)
	is.NoErr(err)
	is.True(r)
}

func Test_Infoln(t *testing.T) {
	is := is.New(t)
	var b bytes.Buffer
	T := &Lumber{
		InfoLog: log.New(&b, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	T.Infoln("test")

	line := b.String()
	line = line[0 : len(line)-1]
	pattern := "^\\[INFO\\].+?test$"
	r, err := regexp.MatchString(pattern, line)
	is.NoErr(err)
	is.True(r)
}

func Test_Infof(t *testing.T) {
	is := is.New(t)
	var b bytes.Buffer
	T := &Lumber{
		InfoLog: log.New(&b, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	T.Infof("%s", "test")

	line := b.String()
	line = line[0 : len(line)-1]
	pattern := "^\\[INFO\\].+?test$"
	r, err := regexp.MatchString(pattern, line)
	is.NoErr(err)
	is.True(r)
}
