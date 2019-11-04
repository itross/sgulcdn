package controller

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/go-chi/chi"
	"github.com/itross/sgul"
	"github.com/itross/sgulcdn/internal/service"
)

var logger = sgul.GetLogger()

// UploadController is the upload endpoint description.
type UploadController struct {
	sgul.Controller
	uploader service.UploadService
}

// BasePath  returns the controller base routing path (implements sgul.RestController).
func (uc *UploadController) BasePath() string {
	return uc.Path
}

// Router return the router for this controller
func (uc *UploadController) Router() chi.Router {
	r := chi.NewRouter()
	r.Get("/", uc.index)
	r.Post("/", uc.upload)
	return r
}

// NewUploadController returns a new UploadController instance.
func NewUploadController(uploader service.UploadService) *UploadController {
	return &UploadController{
		Controller: sgul.NewController("/upload"),
		uploader:   uploader,
	}
}

func (uc *UploadController) upload(w http.ResponseWriter, r *http.Request) {
	requestID := middleware.GetReqID(r.Context())
	logger.Infow("Request to upload file", "request-id", requestID)
	if err := uc.uploader.Upload(r.Context(), r, "file"); err != nil {
		logger.Errorw("error uploading file", "error", err, "request-id", requestID)
		uc.RenderError(w, sgul.NewHTTPError(err, http.StatusInternalServerError, "Unable to upload the file", requestID))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, "Upload OK")
}

func (uc *UploadController) index(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.HTML(w, r, "<html><body><form method=\"POST\" action=\"/cdn/upload\" enctype=\"multipart/form-data\"><input type=\"file\" name=\"file\"><input type=\"submit\"></form></body></html>")
}
