package service

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/itross/sgul"
	"github.com/olahol/go-imageupload"
)

var logger = sgul.GetLogger()

// UploadService describe the interface that provides the upload functionalities.
type UploadService interface {
	Upload(ctx context.Context, r *http.Request, field string) error
}

type uploader struct {
}

// NewUploader returns a new uploader instance.
func NewUploader() UploadService {
	return &uploader{}
}

// Upload will manage uploads calls saving file into the user-id store filesystem.
func (u *uploader) Upload(ctx context.Context, r *http.Request, field string) error {
	img, err := imageupload.Process(r, field)
	if err != nil {
		return err
	}

	thumb, err := imageupload.ThumbnailPNG(img, 250, 250)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("./store/tmp.png", thumb.Data, 0644); err != nil {
		return err
	}

	logger.Infow("file saved", "file", "./store/tmp.png", "request-id", middleware.GetReqID(ctx))
	return nil
}
