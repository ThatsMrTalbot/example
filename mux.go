package example

import (
	"net/http"

	"github.com/go-zoo/bone"
)

// Mux is an injectable mux instance
type Mux struct {
	*bone.Mux
}

// Start initializes mux
func (mux *Mux) Start() error {
	mux.Mux = bone.New()
	return nil
}

// GetAllValues passes through to bone
func GetAllValues(req *http.Request) map[string]string {
	return bone.GetAllValues(req)
}

// GetValue passes through to bone
func GetValue(req *http.Request, key string) string {
	return bone.GetValue(req, key)
}
