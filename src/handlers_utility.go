package conductor

import (
	"fmt"
	"net/http"
)

func utilPingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
