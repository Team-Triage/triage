package filter

import (
	filterChan "triage/channels/toFilter"
)

func filter() {
	for {
		msg := filterChan.GetMessage()
		if msg.status >= 1 {
			// write ack
		} else {
			// send to reaper
			// nack
		}

	}
}
