package consumerManager

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/team-triage/triage/channels/newConsumers"
	"github.com/team-triage/triage/types"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	var consReq types.ConsumerRequest
	fmt.Println(req)
	err := json.NewDecoder(req.Body).Decode(&consReq)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// TODO: need to add auth checking
	address := consReq.Address
	fmt.Printf("CONSUMER MANAGER: Consumer requested connection from: %v\n", address)

	newConsumers.AppendMessage(address)
	w.WriteHeader(200)
}
