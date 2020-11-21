package helpers

import (
	"net/http"

	"github.com/cecepsprd/crowfu-api/internal/model"
)

func GetResponse(code int, data ...interface{}) model.APIResponse {

	switch code {
	case http.StatusOK:
		return model.APIResponse{
			Status: "success",
			Data:   data,
		}
	case http.StatusNotFound:
		return model.APIResponse{
			Status: "Data Not Found :(",
			Data:   data,
		}
	case http.StatusConflict:
		return model.APIResponse{
			Status: "Data already exists",
			Data:   data,
		}
	case http.StatusBadRequest:
		return model.APIResponse{
			Status: "Given param is not valid",
			Data:   data,
		}
	default:
		return model.APIResponse{
			Status: "Internal server error :(",
			Data:   data,
		}
	}

}
