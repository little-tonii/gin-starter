package constant

type contextKey struct {
	CLAIMS       string
	REQUEST_DATA string
}

var ContextKey *contextKey

func InitContextKey() {
	ContextKey = &contextKey{
		CLAIMS:       "claims",
		REQUEST_DATA: "request_data",
	}
}
