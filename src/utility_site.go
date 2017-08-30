package conductor

import (
	"errors"
	"net/http"
)

func GetSite(r *http.Request) (Site, error) {
	var ss Sites
	var s Site
	host := r.URL.Host

	ss.Domain = host
	ss.Get()

	if len(ss.Sites) > 0 {
		s = ss.Sites[0]
	} else {
		return s, errors.New("NO")
	}

	return s, nil
}
