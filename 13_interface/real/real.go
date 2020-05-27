package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

// Retriever ...x
type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

// Get ...x
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}
