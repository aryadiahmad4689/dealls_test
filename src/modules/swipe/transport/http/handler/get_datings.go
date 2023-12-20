package handler

import (
	"net/http"

	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/transport/http/response"
)

func (handler *Handler) GetDating(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		ctx    = r.Context()
		result interface{}
	)

	result, err = handler.endpoint.GetDatings(ctx)
	if err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, "Failed", err.Error())
		return
	}

	// Send the result back as JSON
	response.SendSuccessResponse(w, http.StatusOK, "Success", result)
}
