package stacks

//IStack interface para la pila
type IStack interface {
	Push(v string)
	Pop() (string, error)
}
