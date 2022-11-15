package consumerManager

import (
	"net/http"
	// "fmt"
	// "encoding/json"
	// "log"
)

func StartHttpServer() {
	http.HandleFunc("/consumers", consumerHandler)
	http.HandleFunc("/", healthCallback)
	http.ListenAndServe(":9000", nil)
}
