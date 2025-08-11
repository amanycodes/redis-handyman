package logx

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", log.LstdFlags)
var verbose bool

func Init(v bool)              { verbose = v }
func Infof(f string, a ...any) { logger.Printf(f, a...) }
func Debugf(f string, a ...any) {
	if verbose {
		logger.Printf("DEBUG: "+f, a...)
	}
}
func Errorf(f string, a ...any) { logger.Printf("ERROR: "+f, a...) }
