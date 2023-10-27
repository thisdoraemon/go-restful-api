package exception

import (
	"net/http"
	"restful-api/helper"
	"restful-api/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, error interface{}) {

	internalServerError(writer, request, error)

}

func internalServerError(writer http.ResponseWriter, request *http.Request, error interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   error,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
