package eval

type Expression interface {
	Evaluate(env Environment) float64
	Check(vars map[Variable]bool) error
	String() string // exercise-07.13
}

type Variable string

type literal float64

type unary struct {
	operator rune
	x Expression
}

type binary struct {
	operator rune
	x, y Expression
}

type call struct {
	function string
	arguments []Expression
}

type Environment map[Variable]float64

var numberParameters = map[string]int{
	"pow": 2,
	"sin": 1,
	"cos": 1,
	"sqrt": 1,
}