package consumerManager

import (
	"net/http"
)

func StartServer(callback func(w http.ResponseWriter, req *http.Request)) {
	http.HandleFunc("/consumers", callback)
	http.ListenAndServe(":9000", nil)
}
