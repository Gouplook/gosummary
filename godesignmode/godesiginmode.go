package godesignmode

// 策略模式
type IStrategy interface {
	do(int, int) int
}

type add struct {
}

func (*add) do(a, b int) int {
	return a + b
}

type reduce struct {
}

func (*reduce) do(a, b int) int {
	return a - b
}

type Operator struct {
	strategy IStrategy
	//IStrategy
}

// 设置策略(用接口来接受）
func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}
