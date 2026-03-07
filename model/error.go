package model

import (
	"encoding/json"
	"fmt"
	"time"
)

const format = "2006-01-02 15:04:05"

type XError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Detail     string `json:"detail"`
	Time       string `json:"time"`
	HttpStatus int    `json:"-"`
}

func (e *XError) Error() string {
	if e == nil {
		return fmt.Sprintf("{\"code\":%d,\"message\":\"%s\"}", 500, "internal error")
	}
	byt, err := json.Marshal(&e)
	if err != nil {
		return fmt.Sprintf("{\"code\":%d,\"message\":\"%s\"}", e.Code, err.Error())
	}
	return string(byt)
}

func NewInvalidArg(s string) *XError {
	return &XError{
		Code:       10000,
		Message:    "invalid argument",
		Detail:     fmt.Sprintf("parameter `%s` is invalid", s),
		Time:       time.Now().Format(format),
		HttpStatus: 400,
	}
}

func NewNotFound() *XError {
	return &XError{
		Code:       10001,
		Message:    "address not found",
		Detail:     "check the URL",
		Time:       time.Now().Format(format),
		HttpStatus: 404,
	}
}

func NewBadRequest() *XError {
	return &XError{
		Code:       10002,
		Message:    "address bad request",
		Detail:     "check the URL",
		Time:       time.Now().Format(format),
		HttpStatus: 400,
	}
}
func NewNotImplemented(m string) *XError {
	return &XError{
		Code:       10003,
		Message:    "not implemented",
		Detail:     fmt.Sprintf("method %s is not implemented", m),
		Time:       time.Now().Format(format),
		HttpStatus: 501,
	}
}
func NewMethodNotAllowed(m string) *XError {
	return &XError{

		Code:       10003,
		Message:    "method not allowed",
		Detail:     fmt.Sprintf("http method `%s` is not allowed", m),
		Time:       time.Now().Format(format),
		HttpStatus: 405,
	}
}

func NewSuccess() *XError {
	return &XError{
		Code:       10000,
		Message:    "success",
		Detail:     "transaction was successfully created",
		Time:       time.Now().Format(format),
		HttpStatus: 200,
	}
}
