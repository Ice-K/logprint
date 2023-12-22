package logprint

import (
	"fmt"
	"time"
)

func Error(message string) {
	t := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Printf("[ERROR]\t%s\t%s\n", t, message)
}