package main

import (
	"isbn-9787115531421/src/unit4/chapter20/life/universe"
	"time"
)

func main() {
	u1, u2 := universe.New(), universe.New()
	u1.Seed()
	for i := 0; i < 300; i++ {
		universe.Step(u1, u2)
		u1.Show()
		time.Sleep(time.Second / 30)
		u1, u2 = u2, u1
	}
}