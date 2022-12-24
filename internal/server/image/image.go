package photo

import (
	"net/http"
	spec "yukiko-shop/internal/generated/spec/image"
	"yukiko-shop/pkg/response"
)

func (s Server) PostImages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		s.logger.Warn(err)
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	res, err := s.imageUseCase.UploadImage(ctx, file, *fileHeader)
	if err != nil {
		s.logger.Warn(err)
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	response.JSON(w, http.StatusCreated, res)
}
