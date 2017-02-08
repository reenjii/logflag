# logflag provides a `-log` command line flag to customize default logrus logger

Simply import package and call `logflag.Parse()` after `flag.Parse()`.

This adds a `-log` multivalue flag to your command line with the following possible values:

- debug, info, warn, error, fatal: logging level
- color: force logrus to color output (text formatter)
- nocolor: force logrus NOT to color output (text formatter)
- json : use `logrus.JSONFormatter` to output logs in `JSON` format

Multiple flags can be provided, either csv style or by writing `-log` flag several times.
Values are processed in order, so the last values prevail.

Command line examples:

- `command -log debug,json` set level to debug and format to json
- `command -log nocolor,warn` set level to warn and disable colors
- `command -log info -log color` set level to info and force colors

## Usage

Here is a sample program that uses `logflag`.

```golang
package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/reenjii/logflag"
)

var str string

func init() {
	flag.StringVar(&str, "hello", "world", "custom flag")
}

func main() {
	flag.Parse()
	logflag.Parse() // Call after regular flag.Parse()

	logrus.WithField("hello", str).Info("custom flag")

	logrus.Debug("Debug log")
	logrus.Info("Info log")
	logrus.Warn("Warn log")
	logrus.Error("Error log")
	logrus.Fatal("Fatal log")
}
```

It produces the following help message.

```bash
Usage of ./test:
  -hello string
        custom flag (default "world")
  -log flags
        log flags, several allowed [debug,info,warn,error,fatal,color,nocolor,json]
```

