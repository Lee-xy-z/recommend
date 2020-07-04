package storage

import "github.com/Lee-xy-z/recommend/storage/rcdstore"

// Factory defines an interface for a factory that can create implementations of different storage components.
// Implementations are also encouraged to implement plugin.Configurable interface.
type Factory interface {
	//CreateRcdReader() (rcdstore.Reader, error)

	// CreateRcdWriter creates a rcdstore.Writer
	CreateRcdWriter() (rcdstore.Writer, error)
}
