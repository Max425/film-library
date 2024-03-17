package dto

import (
	"context"
	"github.com/Max425/film-library.git/internal/constants"
	"net/http"
)

type RequestInfo struct {
	Status int
}

type SuccessClientResponseDto struct {
	Result interface{} `json:"result"`
}

type ErrorClientResponseDto struct {
	Error interface{} `json:"error"`
}

func NewSuccessClientResponseDto(ctx context.Context, w http.ResponseWriter, statusCode int, payload interface{}) {
	response := SuccessClientResponseDto{
		Result: payload,
	}

	responseJSON, err := response.MarshalJSON()
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	sendData(ctx, w, statusCode, responseJSON)
}

func NewErrorClientResponseDto(ctx context.Context, w http.ResponseWriter, statusCode int, payload interface{}) {
	response := ErrorClientResponseDto{
		Error: payload,
	}

	responseJSON, err := response.MarshalJSON()
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	sendData(ctx, w, statusCode, responseJSON)
}

func sendData(ctx context.Context, w http.ResponseWriter, statusCode int, responseJSON []byte) {
	requestInfo, ok := ctx.Value(constants.KeyRequestInfo).(*RequestInfo)
	if ok {
		requestInfo.Status = statusCode
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
