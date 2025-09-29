package inwx

import (
	"fmt"
)

type APIError struct {
	Code       int    `json:"code"`
	Message    string `json:"msg"`
	ReasonCode string `json:"reasonCode,omitempty"`
	Reason     string `json:"reason,omitempty"`
}

func (e *APIError) Error() string {
	if e.ReasonCode != "" {
		return fmt.Sprintf("API error %d (%s): %s - %s", e.Code, e.ReasonCode, e.Message, e.Reason)
	}
	return fmt.Sprintf("API error %d: %s", e.Code, e.Message)
}

func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

func NewAPIErrorWithReason(code int, message, reasonCode, reason string) *APIError {
	return &APIError{
		Code:       code,
		Message:    message,
		ReasonCode: reasonCode,
		Reason:     reason,
	}
}
