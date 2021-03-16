/*
Package logflag provides command line flag "-log" to setup default logrus logger.

Simply import package and call
	logflag.Parse()
after
	flag.Parse()

This adds a "-log" multivalue flag to your command line with the following possible values:

	- debug, info, warn, error, fatal: logging level
	- color: force logrus to color output (text formatter)
	- nocolor: force logrus NOT to color output (text formatter)
	- json : use logrus.JSONFormatter to output logs in JSON format

Command line examples:

	command -log debug,json
	command -log nocolor,warn
	command -log info -log color
*/
package logflag

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

// stringslice stores multi-value command line arguments.
type stringslice []string

// String makes stringslice implement flag.Value interface.
func (s *stringslice) String() string {
	return fmt.Sprintf("%s", *s)
}

// Set makes stringslice implement flag.Value interface.
func (s *stringslice) Set(value string) error {
	for _, v := range strings.Split(value, ",") {
		if len(v) > 0 {
			*s = append(*s, v)
		}
	}
	return nil
}

var (
	logflags stringslice
	levels   = regexp.MustCompile("^(debug|info|warn|error|fatal)$")
	colors   = regexp.MustCompile("^(no)?colou?rs?$")
	json     = regexp.MustCompile("^json$")
)

func init() {
	flag.Var(&logflags, "log", "log `flags`, several allowed [debug,info,warn,error,fatal,color,nocolor,json]")
}

// Parse parses the -log flags and initializes the default logrus logger.
// It should be called after regular flag.Parse() call.
func Parse() {
	for _, f := range logflags {
		if levels.MatchString(f) {
			lvl, err := logrus.ParseLevel(f)
			if err != nil {
				// Should never happen since we select correct levels
				// Unless logrus commits a breaking change on level names
				panic(fmt.Errorf("invalid log level: %s", err.Error()))
			}
			logrus.SetLevel(lvl)
		} else if colors.MatchString(f) {
			if f[:2] == "no" {
				logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
			} else {
				logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
			}
		} else if json.MatchString(f) {
			logrus.SetFormatter(&logrus.JSONFormatter{})
		}
	}
}
