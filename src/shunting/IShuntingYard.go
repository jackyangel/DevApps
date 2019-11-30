package shunting

type IShuntingYard interface {
	Evaluate(infix string) (postfix string)
}
