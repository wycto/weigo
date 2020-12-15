package weigo

/*
响应类，响应的处理
*/
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
