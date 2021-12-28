package restfulnessapplication

import (
	"fmt"
	"net/http"
)

type HttpHandler struct {
}

func (httpHandler *HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	println(req.URL.Path)
	_, _ = fmt.Fprintf(w, "Hi")

}
