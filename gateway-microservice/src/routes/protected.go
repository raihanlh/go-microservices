package routes

type HTTPRequest struct {
	Method string
	Path   string
}

func NewHTTPRequest(method string, path string) HTTPRequest {
	return HTTPRequest{
		Method: method,
		Path:   path, // Note: Use regular expression to define path
	}
}

var protected = []HTTPRequest{
	NewHTTPRequest("GET", "\\/user"),
	NewHTTPRequest("POST", "\\/articles\\/\\w+"),
	NewHTTPRequest("PATCH", "\\/articles\\/\\w+"),
}
