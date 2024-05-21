package response

type RecoverableExceptionResponse struct {
	RecoverableExceptionType int32 `json:"recoverable_exception_type" enum:"RecoverableExceptionType"`
}
