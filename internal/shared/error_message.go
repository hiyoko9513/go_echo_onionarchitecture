package shared

import "net/http"

func GetErrorMessage(statusCode int) string {
	switch statusCode {
	case http.StatusBadRequest:
		return "Bad request"
	case http.StatusUnauthorized:
		return "Unauthorized"
	case http.StatusForbidden:
		return "Forbidden"
	case http.StatusNotFound:
		return "Not found"
	case http.StatusInternalServerError:
		return "Internal server error"
	default:
		return "Error"
	}
}
