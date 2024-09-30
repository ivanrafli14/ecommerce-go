package response

import "github.com/gin-gonic/gin"

type Response struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	ErrorMessage string `json:"error_message,omitempty"`
	ErrorCode    string `json:"error_code,omitempty"`
	Payload      any    `json:"payload,omitempty"`
	Pagination   any    `json:"pagination,omitempty"`
}

func Success(c *gin.Context, httpCode int, message string, payload any) {
	c.JSON(httpCode, Response{
		Success: true,
		Message: message,
		Payload: payload,
	})
}

func SuccessWithPagination(c *gin.Context, httpCode int, message string, payload any, pagination any) {
	c.JSON(httpCode, Response{
		Success:    true,
		Message:    message,
		Payload:    payload,
		Pagination: pagination,
	})
}

func Failed(c *gin.Context, Error error, message string) {
	myErr, ok := ErrorMapping[Error.Error()]
	if !ok {
		myErr = ErrorGeneral
	}

	c.JSON(myErr.HttpCode, Response{
		Success:      false,
		Message:      myErr.Message,
		ErrorMessage: message,
		ErrorCode:    myErr.ErrorCode,
	})
}
