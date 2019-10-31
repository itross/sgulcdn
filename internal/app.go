package internal

import e "github.com/itross/sgulengine"

// New returns a new sgul Engine instance for the CDN service app.
func New() *e.Engine {
	return e.NewWith(e.NewDefaultAPIComponent())
}
