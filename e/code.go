package e

const (
	OK            = 200 // OK : The server successfully accepts the request from the client
	ERROR         = 500 // ERROR : The server has an unknown e
	InvalidParams = 400 // InvalidParams : The client's request has illegal parameters
	Unauthorized  = 401 // Unauthorized 请求要求身份验证
	Forbidden     = 403 // Forbidden The client does not have access rights to the content
	NotFound      = 404 // NotFound : The server can not find the requested resource
)

var Err = New(ERROR)

var ErrInvalidParams = New(InvalidParams)

var ErrUnauthorized = New(Unauthorized)

var ErrForbidden = New(Forbidden)

var ErrNotFound = New(NotFound)
