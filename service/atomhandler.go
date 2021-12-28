package service

type AtomHandler interface {
	DoMapping(atomMapper *AtomMapper)
	DoServe(atom *Atom, req *Request) *Response
}
