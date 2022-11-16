package consumerManager

import (
	"net/http"
)

func StartHttpServer() {
	http.HandleFunc("/consumers", consumerHandler)
	http.HandleFunc("/", healthCallback)
	http.ListenAndServe(":9000", nil)
}
