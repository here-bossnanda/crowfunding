package helper

type response struct {
	meta meta
	data interface{}
}

type meta struct {
	message string
	code    int
	status  string
}
