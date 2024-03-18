package dto

import (
	"context"
	"github.com/Max425/film-library.git/internal/common/constants"
	"log"
	"net/http"
)

type RequestInfo struct {
	Status  int
	Message string
}

type ClientResponseDto struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Payload any    `json:"payload"`
}

func NewSuccessClientResponseDto(ctx context.Context, w http.ResponseWriter, payload any) {
	response := ClientResponseDto{
		Status:  http.StatusOK,
		Message: "success",
		Payload: payload,
	}
	sendData(ctx, w, response, http.StatusOK, "success")
}

func NewErrorClientResponseDto(ctx context.Context, w http.ResponseWriter, statusCode int, message string) {
	response := ClientResponseDto{
		Status:  statusCode,
		Message: message,
		Payload: "",
	}
	sendData(ctx, w, response, statusCode, message)
}

func sendData(ctx context.Context, w http.ResponseWriter, response ClientResponseDto, statusCode int, message string) {
	responseJSON, err := response.MarshalJSON()
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	requestInfo, ok := ctx.Value(constants.KeyRequestInfo).(*RequestInfo)
	if !ok {
		log.Println("Request info not found in context")
	} else {
		requestInfo.Status = statusCode
		requestInfo.Message = message
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
