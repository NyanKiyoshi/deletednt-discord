package deletednt

import (
	"net/http"
	"time"
)

func getHttpClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}
