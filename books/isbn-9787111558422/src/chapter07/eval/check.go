package eval

import (
	"fmt"
	"strings"
)

func (v Variable) Check(vars map[Variable]bool) error {
	vars[v] = true
	return nil
}

func (l literal) Check(vars map[Variable]bool) error {
	return nil
}

func (u unary) Check(vars map[Variable]bool) error {
	if !strings.ContainsRune("+-", u.operator) {
		return fmt.Errorf("unexpected unary operator %q", u.operator)
	}
	return u.x.Check(vars)
}

func (b binary) Check(vars map[Variable]bool) error {
	if !strings.ContainsRune("+-*/", b.operator) {
		return fmt.Errorf("unexpected unary operator %q", b.operator)
	}
	err1, err2 := b.x.Check(vars), b.y.Check(vars)
	if err1 != nil {
		return err1
	}
	return err2
}

func (c call) Check(vars map[Variable]bool) error {
	arity, ok := numberParameters[c.function]
	if !ok {
		return fmt.Errorf("unknown function %q", c.function)
	}
	if len(c.arguments) != arity {
		return fmt.Errorf("call to %s has %d arguments, want %d", c.function, len(c.arguments), arity)
	}
	for _, arg := range c.arguments {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}