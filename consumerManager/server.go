package consumerManager

import (
	"log"
	"net/http"
)

func StartHttpServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/consumers", consumerHandler)
	mux.HandleFunc("/", healthCallback)
	err := http.ListenAndServe(":9000", mux)

	if err != nil {
		log.Println(err)
	}
}
