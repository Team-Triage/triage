package consumerManager

import (
	"net/http"
	// "fmt"
	// "encoding/json"
	// "log"
)


func healthCallback(w http.ResponseWriter, req *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("200"))
}


func StartServer(callback func(w http.ResponseWriter, req *http.Request)) {
	http.HandleFunc("/consumers", callback)
	http.HandleFunc("/health", healthCallback)
	http.ListenAndServe(":9000", nil)
}
