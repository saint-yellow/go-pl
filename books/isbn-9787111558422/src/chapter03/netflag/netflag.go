package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBoardcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool {
	return v & FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBoardcast(v *Flags) {
	*v |= FlagBoardcast
}

func IsCast(v Flags) bool {
	return v & (FlagBoardcast | FlagMulticast) != 0
}

func main() {
	v := FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))

	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))

	SetBoardcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}