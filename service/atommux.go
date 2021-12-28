package Service

type AtomMux struct {
	handlerMap map[string]AtomHandler
}

func NewAtomMux() *AtomMux {
	return &AtomMux{}
}

func (atomMux *AtomMux) Handle(pattern string, atomHandler AtomHandler) {
	atomMux.panicIfHandlerExist(pattern)
}

func (atomMux *AtomMux) HandleFunc(pattern string, handler func(servAtom *Atom, req *Request) *Response) {
	atomMux.panicIfHandlerExist(pattern)
}

func (atomMux *AtomMux) panicIfHandlerExist(pattern string) {
	_, existed := atomMux.handlerMap[pattern]
	if existed {
		Panic("Handler for this pattern is already existed")
	}
}
