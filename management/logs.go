package management

import (
	"log"
	"os"
)

var (
	// Info writes logs in the color blue with "INFO: " as prefix
	Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)

	// Warning writes logs in the color yellow with "WARNING: " as prefix
	Warning = log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)

	// Error writes logs in the color red with "ERROR: " as prefix
	Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)

	// Debug writes logs in the color cyan with "DEBUG: " as prefix
	Debug = log.New(os.Stdout, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)
)