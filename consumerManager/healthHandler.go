package consumerManager

import (
	"fmt"
	"net/http"
)

func healthCallback(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Health check received!")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200"))
}
