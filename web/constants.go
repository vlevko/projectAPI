package web

const (
	// Errors
	responseError         = "error"
	badRequestError       = "bad request"
	pageNotFoundError     = "page not found"
	projectNotFoundError  = "project not found"
	columnNotFoundError   = "column not found"
	taskNotFoundError     = "task not found"
	commentNotFoundError  = "comment not found"
	columnNotDeletedError = "column cannot be deleted if exists"
	defaultError          = "something wrong"
	// Success
	responseSuccess        = "success"
	projectDeletedResponse = "project deleted"
	columnDeletedResponse  = "column deleted"
	taskDeletedResponse    = "task deleted if existed"
	commentDeletedResponse = "comment deleted"
	columnPositionResponse = "column position changed"
	taskPositionResponse   = "task position changed"
	taskStatusResponse     = "task status changed"
	// Listening port
	port = "8000"
)
