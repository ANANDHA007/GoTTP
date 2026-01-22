package http

type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
	Close   bool
}

type Response struct {
	Status  string
	Headers map[string]string
	Body    []byte
	Close   bool
}
