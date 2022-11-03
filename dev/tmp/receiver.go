package tmp

import (
	"fmt"
	dispatchChan "triage/channels/toDispatch"
	"triage/data/commitTable"
)

func Receiver() {
	for {
		msg := dispatchChan.GetMessage()
		fmt.Println(msg.Value)

		fmt.Println(commitTable.CommitHash)
	}
}
