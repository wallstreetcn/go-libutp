package utp

import (
	"log"
	"os"
)

var (
	logCallbacks = false
	utpLogging   = false
)

var Logger = log.New(os.Stderr, "go-libutp: ", log.LstdFlags|log.Lshortfile)

// SetUTPLogging sets the utp logging open
func SetUTPLogging(open bool) {
	utpLogging = open
}

// SetUTPLogging sets the callback logging open
func SetLogCallBacks(open bool) {
	logCallbacks = open
}
