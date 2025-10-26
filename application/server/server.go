package application

import (
	"fmt"
	"net/http"
)

func PlayerServer(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "20")
}
