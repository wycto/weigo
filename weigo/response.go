package weigo

import (
	"net/http"
	"time"
)

type Response struct {
	http.ResponseWriter
	Started bool
	Status  int
	Elapsed time.Duration
}
