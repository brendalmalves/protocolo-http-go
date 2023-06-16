type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    string
}

type Response struct {
	Status  int
	Headers map[string]string
	Body    string
}
