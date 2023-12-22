package logprint

import (
	"fmt"
	"time"
)

func Debug(message interface{}) {
	t := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Printf("[DEBUG]\t%s\t%s\n", t, message)
}
