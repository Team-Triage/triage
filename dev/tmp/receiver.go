package tmp

import (
	"fmt"
	"time"

	"github.com/team-triage/triage/data/commitTable"
)

func Receiver() {
	for {
		fmt.Println(commitTable.CommitHash)
		time.Sleep(time.Millisecond * 200)
	}
}
