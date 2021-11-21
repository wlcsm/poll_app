package server

// Very light wrapper over the http server in net/http
import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wlcsm/poll_app/common"
)

type HandlerFunc = func(r *http.Request) common.HTTPResponse

type HTTPServer struct {
	port int
}

func New(port int) HTTPServer {
	return HTTPServer{
		port: port,
	}
}

func (h *HTTPServer) Start() error {
	return http.ListenAndServe(":"+fmt.Sprint(h.port), http.DefaultServeMux)
}

// This doesn't support multiple handlers on the same path
func (h *HTTPServer) AddRoute(method, path string, handler HandlerFunc) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			return
		}

		res := handler(r)
		bytes, err := json.Marshal(res)
		if err != nil {
			return
		}

		w.Write(bytes)
	})
}
