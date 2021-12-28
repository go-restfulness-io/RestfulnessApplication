package service

type Stage uint64

const (
	UNKNOWN Stage = iota
	ERROR
	NEW
	MAPPING
	SERVE
)

type Method string

const (
	MethodGet     Method = "GET"
	MethodHead           = "HEAD"
	MethodPost           = "POST"
	MethodPut            = "PUT"
	MethodPatch          = "PATCH"
	MethodDelete         = "DELETE"
	MethodConnect        = "CONNECT"
	MethodOptions        = "OPTIONS"
	MethodTrace          = "TRACE"
)

type Path string
