package consumerManager

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/team-triage/triage/channels/newConsumers"
)

func handler(w http.ResponseWriter, req *http.Request) {
	dummyToken := "dummyToken"

	if entry, ok := req.Header["Authorization"]; ok {
		if dummyToken != entry[0] { // token doesn't match
			http.Error(w, "Malformed or invalid authorization token", 401)
			return
		}
	} else { // missing token
		http.Error(w, "Missing Authorization header", 400)
		return
	}

	remoteIp := strings.Split(req.RemoteAddr, ":")[0]

	var grpcPort string

	if entry, ok := req.Header["Grpcport"]; ok {
		grpcPort = entry[0]
	} else { // no grpcport in header
		http.Error(w, "Missing gRPC port header", 401)
		return
	}

	consumerAddress := remoteIp + ":" + grpcPort
	fmt.Printf("CONSUMER MANAGER: Consumer requested connection from: %v\n", consumerAddress)

	newConsumers.AppendMessage(consumerAddress)
	w.WriteHeader(200)
}
