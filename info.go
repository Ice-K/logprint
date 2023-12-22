package logprint

import (
	"fmt"
	"time"
)

func Info(message interface{}) {
	t := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Printf("[INFO]\t%s\t%s\n", t, message)
}