package utils

import (
	"log"
	"os"
)

var Log = log.New(os.Stdout, "[go-search] ", log.Lshortfile|log.Ldate|log.Ltime)
