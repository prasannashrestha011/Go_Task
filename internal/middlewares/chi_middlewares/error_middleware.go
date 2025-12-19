package chimiddlewares

import (
	"main/internal/schema"
	"main/internal/utils"
	"net/http"
)

type errContextKey string



type errorResponseWriter struct {
    http.ResponseWriter
    err *utils.AppError
}


func SetError(w http.ResponseWriter, err *utils.AppError) {
    if ew, ok := w.(*errorResponseWriter); ok {
        ew.err = err
    }
}

func ErrorMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        wrapped := &errorResponseWriter{ResponseWriter: w}
        
        defer func() {
            if rec := recover(); rec != nil {
                resp := schema.ErrorResponse("INTERNAL_SERVER_ERR", "Internal server error", "")
                utils.JsonResponseWriter(w, http.StatusInternalServerError, resp)
            }
        }()
        
        next.ServeHTTP(wrapped, r)  // ← Pass wrapped writer
        
        if wrapped.err != nil {
            resp := schema.ErrorResponse(wrapped.err.Code, wrapped.err.Message, wrapped.err.Details)
            utils.JsonResponseWriter(w, wrapped.err.StatusCode, resp)
            return
        }
    })
}