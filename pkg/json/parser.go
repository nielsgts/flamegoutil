package parser

import (
	"net/http"
	"github.com/flamego/flamego"
	"encoding/json"
)

// ParseBody is a Flamego handler that parses the JSON body of an HTTP request
// into a struct of type I and maps it to the context.
// If parsing fails, it responds with a 400 Bad Request error.
func ParseBody[I interface{}]() flamego.Handler {
	return flamego.ContextInvoker(func(c flamego.Context) {
		var data I

		if err := json.NewDecoder(c.Request().Body().ReadCloser()).Decode(&data); err != nil {
			http.Error(c.ResponseWriter(), "invalid JSON: " + err.Error(), http.StatusBadRequest)
			return
		}
		c.Map(data)
	})
}