package surra

import (
	"github.com/surra/internal/json"
)

type TMFError struct {
	Code    string `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func (t *TMFError) Set(code string, reason string, message string) []byte {
	def := TMFError{Code: code, Reason: reason, Message: message}
	res, _ := json.Marshal(def)
	return res
}

var defaultErr = TMFError{}

var (
	default404Body = defaultErr.Set("ERR-001", "Not Found", "Resource Not Found")
	default405Body = defaultErr.Set("ERR-002", "Method Not Allowed", "The request method is not allowed")
)

//
//func OnJsonBindError(c *Context, err error) {
//	var jsonErr *json.UnmarshalTypeError
//	var ver validator.ValidationErrors
//
//	if errors.As(err, &ver) {
//
//		c.JSON(http.StatusBadRequest, dynamicErrors.Fields(ver))
//		return
//	}
//
//	if errors.As(err, &jsonErr) {
//
//		c.JSON(400, err.Error())
//		return
//	}
//	if err.Error() == "EOF" || err.Error() == "unexpected EOF" {
//		c.AbortWithStatusJSON(http.StatusBadRequest, staticErrors.InvalidBody())
//		return
//	}
//	c.AbortWithStatusJSON(http.StatusInternalServerError, staticErrors.InternalServerError())
//	return
//}
