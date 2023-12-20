package handler

import (
	"net/http"

	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/transport/http/response"
)

func (handler *Handler) GetPackage(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		ctx    = r.Context()
		result interface{}
	)

	result, err = handler.endpoint.GetPackage(ctx)
	if err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, "Failed", err.Error())
		return
	}

	// Send the result back as JSON
	response.SendSuccessResponse(w, http.StatusOK, "Success", result)
}
