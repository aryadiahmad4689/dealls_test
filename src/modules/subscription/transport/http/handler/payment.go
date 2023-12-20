package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/endpoint"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/transport/http/response"
)

func (handler *Handler) Payment(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		data   endpoint.PaymentRequest
		ctx    = r.Context()
		result interface{}
	)
	// Decode the request body into data
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, "Failed SignIn", err.Error())
		return
	}
	result, err = handler.endpoint.Payment(ctx, data)
	if err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, "Failed", err.Error())
		return
	}

	// Send the result back as JSON
	response.SendSuccessResponse(w, http.StatusOK, "Success", result)
}
