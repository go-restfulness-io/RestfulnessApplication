package Service

import (
	"fmt"
	"reflect"
)

type AtomRegistry map[Atom]AtomHandler

var (
	atomRegistry AtomRegistry = make(map[Atom]AtomHandler)
)

type Atom struct {
	stage    Stage
	request  Request
	response Response
}

type AtomMapper struct {
	name        string
	atom        Atom
	atomHandler AtomHandler
}

func NewAtomMapper(atom Atom, atomHandler AtomHandler) *AtomMapper {
	name := reflect.TypeOf(atomHandler).String()
	println("New Atom Mapper")
	return &AtomMapper{name: name, atom: atom, atomHandler: atomHandler}
}

func NewAtom(atomHandler AtomHandler) *Atom {
	atom := Atom{
		stage: NEW,
	}
	atomRegistry[atom] = atomHandler
	println("New atom")
	return &atom
}

func GetAtomSet() []Atom {
	var atomSet []Atom

	for atom, _ := range atomRegistry {
		atomSet = append(atomSet, atom)
	}

	return atomSet
}

func GetAtomRegistry() AtomRegistry {

	return atomRegistry
}

func GetAtomHandler(atom Atom) AtomHandler {
	return atomRegistry[atom]
}

func (atomMapper *AtomMapper) RequestMapping(method Method, path Path) *Response {
	fmt.Printf("Mapping %s:%s -> %s", method, path, atomMapper.name)

	return &atomMapper.atom.response
}

func (atomMapper *AtomMapper) MethodGetMapping(path Path) *Response {
	return atomMapper.RequestMapping(MethodGet, path)
}

func (atomMapper *AtomMapper) MethodHeadMapping(path Path) *Response {
	return atomMapper.RequestMapping(MethodHead, path)
}

func (atomMapper *AtomMapper) MethodPostMapping(path Path) *Response {
	return atomMapper.RequestMapping(MethodPost, path)
}

func (atomMapper *AtomMapper) MethodPutMapping(path Path) *Response {
	return atomMapper.RequestMapping(MethodPut, path)
}

func (atomMapper *AtomMapper) MethodPatchMapping(path Path) *Response {
	return atomMapper.RequestMapping(MethodPatch, path)
}

func (atomMapper *AtomMapper) MethodDeleteMapping(path Path) *Response {
	return atomMapper.RequestMapping(MethodDelete, path)
}

func (atomMapper *AtomMapper) MethodConnectMapping(path Path) *Response {
	return atomMapper.RequestMapping(MethodConnect, path)
}

func (atomMapper *AtomMapper) MethodOptionsMapping(path Path) *Response {
	return atomMapper.RequestMapping(MethodOptions, path)
}

func (atomMapper *AtomMapper) MethodTraceMapping(path Path) *Response {
	return atomMapper.RequestMapping(MethodTrace, path)
}

func (atomMapper *AtomMapper) BindAll(path Path) *Response {
	println("Binding Service to", path)
	return &atomMapper.atom.response
}
func (atom *Atom) GetStage() Stage {
	return atom.stage
}
func (atom *Atom) SetStage(stage Stage) {
	atom.stage = stage
}

func (atom *Atom) ErrorResponse(message string) *Response {
	preservedStage := atom.stage
	atom.stage = ERROR
	return atom.response.Error(preservedStage, message)
}
