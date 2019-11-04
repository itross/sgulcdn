package controller

import (
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/itross/sgul"
)

// DownloadController is the download endpoint description.
type DownloadController struct {
	sgul.Controller
	// downloader service.DownloadController
}

// BasePath returns the controller base routing path (implements sgul.RestController).
func (dc *DownloadController) BasePath() string {
	return dc.Path
}

// Router return the router for this controller
func (dc *DownloadController) Router() chi.Router {
	r := chi.NewRouter()
	r.Get("/{filename}", dc.download)
	return r
}

// NewDownloadController returns a new DownloadController instance.
func NewDownloadController() *DownloadController {
	return &DownloadController{
		Controller: sgul.NewController("/files/img"),
	}
}

func (dc *DownloadController) download(w http.ResponseWriter, r *http.Request) {
	requestID := middleware.GetReqID(r.Context())
	filename := chi.URLParam(r, "filename")
	logger.Infow("Request to download file", "file", filename, "request-id", requestID)

	f, err := os.Open("./store/tmp.png")
	defer f.Close()
	if err != nil {
		logger.Errorw("error downloading file", "error", err, "request-id", requestID)
		dc.RenderError(w, sgul.NewHTTPError(err, http.StatusInternalServerError, "Unable to download the file", requestID))
		return
	}

	io.Copy(w, f)
}
