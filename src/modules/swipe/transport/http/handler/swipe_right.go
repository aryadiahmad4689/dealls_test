package handler

import (
	"net/http"

	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/endpoint"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/transport/http/response"
	"github.com/go-chi/chi"
)

func (handler *Handler) SwipeRight(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		req    endpoint.SwipeRightReq
		ctx    = r.Context()
		result interface{}
	)
	// Mengambil userID dari URL
	userId := chi.URLParam(r, "user_id")
	req.IsSwipeUserId = userId
	result, err = handler.endpoint.SwipeRight(ctx, req)
	if err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, "Failed", err.Error())
		return
	}

	// Send the result back as JSON
	response.SendSuccessResponse(w, http.StatusOK, "Success", result)
}
