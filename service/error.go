package service

import "errors"

type Error struct {
	stage Stage
	error error
}

func newError(stage Stage, message string) *Error {
	return &Error{stage: stage, error: errors.New("Package Service: " + message)}
}

func Panic(message string) {
	//suppress intellij bug on type any
	panic(any("Package Service: " + message))

}
