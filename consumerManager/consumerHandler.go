package consumerManager

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/team-triage/triage/channels/newConsumers"
)

var authenticationToken string

func SetToken(token string) {
	authenticationToken = token
}

func consumerHandler(w http.ResponseWriter, req *http.Request) {
	if entry, ok := req.Header["Authorization"]; ok {
		if authenticationToken != entry[0] { // token doesn't match
			http.Error(w, "Your token doesn't match; we expected:"+authenticationToken+"\n"+"We got:"+entry[0], 401)
			// http.Error(w, "Malformed or invalid authorization token", 401)
			return
		}
	} else { // missing token
		http.Error(w, "Missing Authorization header", 400)
		return
	}

	var grpcPort string

	if entry, ok := req.Header["Grpcport"]; ok {
		grpcPort = entry[0]
	} else { // no grpcport in header
		http.Error(w, "Missing gRPC port header", 400)
		return
	}

	var remoteIp string

	if entry, ok := req.Header["X-Forwarded-For"]; ok {
		fmt.Printf("HTTP SERVER: Found x-forwarded-for header: %v\n", entry)
		remoteIp = strings.Split(entry[0], ":")[0]
	} else { // should remove this clause for deployment, since it's only for localized testing and triage will always be deployed behind a load balancer
		fmt.Println("HTTP SERVER: Did not find x-forwarded-for header!")
		remoteIp = strings.Split(req.RemoteAddr, ":")[0]
	}

	consumerAddress := remoteIp + ":" + grpcPort
	fmt.Printf("CONSUMER MANAGER: Consumer requested connection from: %v\n", remoteIp)

	newConsumers.AppendMessage(consumerAddress)
	fmt.Fprintf(w, "Prepare to receive messages!")
}
