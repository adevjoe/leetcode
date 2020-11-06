package common

type Stack interface {
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	Len() int
}
