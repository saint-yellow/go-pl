package eval

import (
	"fmt"
	"strings"
)

// exercise-07.13
func (v Variable) String() string {
	return string(v)
}

// exercise-07.13
func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

// exercise-07.13
func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.operator, u.x)
}

// exercise-07.13
func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x, b.operator, b.y)
}

// exercise-07.13
func (c call) String() string {
	var args []string
	for _, a := range c.arguments {
		args = append(args, a.String())
	}
	return fmt.Sprintf("%s(%s)", c.function, strings.Join(args, ", "))
}