package eval

import (
	"fmt"
	"math"
)

func (v Variable) Evaluate(env Environment) float64 {
	return env[v]
}

func (l literal) Evaluate(_ Environment) float64 {
	return float64(l)
}

func (u unary) Evaluate(env Environment) float64 {
	switch u.operator {
	case '+':
		return +u.x.Evaluate(env)
	case '-':
		return -u.x.Evaluate(env)
	}
	panic(fmt.Sprintf("unsupported unary operator； %q", u.operator))
}

func (b binary) Evaluate(env Environment) float64 {
	x, y := b.x.Evaluate(env), b.y.Evaluate(env)
	switch b.operator {
	case '+':
		return x+y
	case '-':
		return x-y
	case '*':
		return x*y
	case '/':
		return x/y
	}
	panic(fmt.Sprintf("unsupported binnary operator； %q", b.operator))
}

func (c call) Evaluate(env Environment) float64 {
	switch c.function {
	case "pow":
		return math.Pow(c.arguments[0].Evaluate(env), c.arguments[1].Evaluate(env))
	case "sin":
		return math.Sin(c.arguments[0].Evaluate(env))
	case "cos":
		return math.Cos(c.arguments[0].Evaluate(env))
	case "sqrt":
		return math.Sqrt(c.arguments[0].Evaluate(env))
	}
	panic(fmt.Sprintf("unsupported function %q", c.function))
}