package internal

import (
	"github.com/itross/sgulcdn/internal/controller"
	"github.com/itross/sgulcdn/internal/service"
	e "github.com/itross/sgulengine"
)

// New returns a new sgul Engine instance for the CDN service app.
func New() *e.Engine {
	uploader := service.NewUploader()
	uc := controller.NewUploadController(uploader)
	dc := controller.NewDownloadController()

	return e.NewWith(e.NewDefaultAPIComponentWith(uc, dc))
}
