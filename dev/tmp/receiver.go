package tmp

import (
	"fmt"
	"time"
	"triage/data/commitTable"
)

func Receiver() {
	for {
		fmt.Println(commitTable.CommitHash)
		time.Sleep(time.Millisecond * 200)
	}
}
