package eval

import (
	"fmt"
	"math"
	"testing"
)

type testCase struct {
	expression string
	environment Environment
	want string
}

func TestEval(t *testing.T) {
	tests := []testCase{
		{"sqrt(A / pi)", Environment{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Environment{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Environment{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Environment{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Environment{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Environment{"F": 212}, "100"},
		// {"x % 2", Environment{"x": 5}, "2"},
		// {"sqrt(x, y)", Environment{"x": 10, "y": 1}, "1"},
	}
	var previousExpression string
	for _, test := range tests {
		if test.expression != previousExpression {
			fmt.Printf("\n%s\n", test.expression)
			previousExpression = test.expression
		}

		expr, err := Parse(test.expression)
		if err != nil {
			t.Error(err)
			continue
		}

		got := fmt.Sprintf("%.6g", expr.Evaluate(test.environment))
		fmt.Printf("\t%v => %s\n", test.environment, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n", test.expression, test.environment, got, test.want)
		}
	}
}
