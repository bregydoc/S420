package s420

import "net/http"

type Plugin interface {
	ProcessFile(r *http.Request, cType ContentType, in []byte) []byte
}
