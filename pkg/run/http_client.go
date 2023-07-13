package run

import "io"

type HttpClient interface {
	NewRequest(method, url string, bodyData io.Reader) ([]byte, error)
}
