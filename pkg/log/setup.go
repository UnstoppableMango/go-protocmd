package log

import (
	"io"
	"os"

	"charm.land/log/v2"
)

func WriteTo(w io.Writer) {
	l := log.New(w)
	if _, ok := os.LookupEnv("DEBUG"); ok {
		l.SetLevel(log.DebugLevel)
	}
	log.SetDefault(l)
}
