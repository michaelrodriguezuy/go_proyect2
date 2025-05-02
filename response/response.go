package response

type Response interface {
	StatusCode() int
	Body() ([]byte, error)
	Error() string
	GetData() any
}
