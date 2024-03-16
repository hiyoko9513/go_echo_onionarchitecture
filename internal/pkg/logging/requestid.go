package logging

import "context"

const RequestIDKey = "X-Request-ID"

func GetRequestIDFromContext(ctx context.Context) string {
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return "none-request-id"
}
