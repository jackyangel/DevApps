package shunting

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/devApps/src/collections/queues"
	"github.com/devApps/src/collections/stacks"
	"github.com/pruebadev/colas"
)

type ShuntingImp struct {
	Operators string
	Stack     stacks.IStack
	Queue     queues.IQueue
}

func (s ShuntingImp) Evaluate(infix string) (postfix string) {

	var op string

	slice := strings.Split(infix, "")
	fmt.Println(slice)

	OpPila := colas.NewStack()
	OpSalida := colas.Init()

	for _, sl := range slice {

		if ValidateNum(sl) {
			OpSalida.Enqueue(sl)
		}

		if sl == "." {
			OpSalida.Enqueue(sl)
		}

		if ValidateOp(sl) {

			if change(sl, op) {
				OpPila.Push(sl)
			} else {
				pop, _ := OpPila.Pop()
				OpPila.Push(sl)
				OpSalida.Enqueue(pop)
			}
			op = sl
		}
	}

	capPila := OpPila.Count

	var pila string
	for num := 0; num < capPila; num++ {
		pila, _ = OpPila.Pop()
		OpSalida.Enqueue(pila)
	}

	postfix = ""

	var pp string

	j := OpSalida.Count

	for i := 0; i < j; i++ {
		pp = *OpSalida.Dequeue()
		postfix = fmt.Sprintf("%s%s", postfix, pp)
		fmt.Println(postfix)
	}

	fmt.Println(postfix)

	return
}

func ValidateNum(num string) bool {
	Re := regexp.MustCompile(`[0-9]+`)
	return Re.MatchString(num)
}

func ValidateOp(op string) bool {
	Re := regexp.MustCompile(`\*|\/|\%|\+|\-`)
	return Re.MatchString(op)
}

func change(push string, pop string) bool {
	change := push > pop
	return change
}
