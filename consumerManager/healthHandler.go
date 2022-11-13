package consumerManager

import "net/http"

func healthCallback(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200"))
}
