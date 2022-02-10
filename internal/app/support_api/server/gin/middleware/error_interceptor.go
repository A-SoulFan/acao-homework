package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/A-SoulFan/acao-homework/internal/pkg/apperrors"
	"github.com/A-SoulFan/acao-homework/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ErrorInterceptor struct {
	logger *zap.Logger
}

func NewErrorInterceptor(logger *zap.Logger) *ErrorInterceptor {
	return &ErrorInterceptor{logger: logger}
}

func (m *ErrorInterceptor) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if len(ctx.Errors) > 0 {
				m.errorResponseHandler(ctx)
				headers, _ := json.Marshal(ctx.Request.Header)
				logs := []zap.Field{
					zap.String("request.method", ctx.Request.Method),
					zap.String("request.url", ctx.Request.URL.String()),
					zap.ByteString("request.headers", headers),
					zap.String("errors", ginErrorsToString(ctx.Errors)),
				}
				m.logger.Error("request error:", logs...)
			}
		}()
		ctx.Next()
	}
}

func (m *ErrorInterceptor) errorResponseHandler(ctx *gin.Context) {
	var resp = &response.JsonResponse{
		Code:    -1,
		Message: "服务器异常，请稍后再试",
		Data:    nil,
	}
	for i := len(ctx.Errors) - 1; i >= 0; i-- {
		err := ctx.Errors[i]
		if appError, ok := errors.Cause(err.Err).(*apperrors.Error); ok {
			switch appError.ErrorType() {
			case apperrors.ValidateError:
				resp.Message = appError.Message()
				ctx.JSON(http.StatusBadRequest, resp)
			case apperrors.AuthenticationError:
				resp.Message = appError.Message()
				ctx.JSON(http.StatusUnauthorized, resp)
			case apperrors.ServiceError:
				ctx.JSON(http.StatusInternalServerError, resp)
			}
			return
		}
	}

	ctx.JSON(http.StatusInternalServerError, resp)
}

func ginErrorsToString(errs []*gin.Error) string {
	if len(errs) == 0 {
		return ""
	}
	var buffer strings.Builder
	for i, msg := range errs {
		_, _ = fmt.Fprintf(&buffer, "Error #%02d: %+v\n", i+1, msg.Err)
		if msg.Meta != nil {
			_, _ = fmt.Fprintf(&buffer, "     Meta: %v\n", msg.Meta)
		}
	}
	return buffer.String()
}
