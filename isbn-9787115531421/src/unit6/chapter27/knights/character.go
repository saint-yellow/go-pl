package main

import "fmt"

type item struct {
	name string
}

type character struct {
	name string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v picks up a %v.\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}

	if c.leftHand == nil {
		fmt.Printf("%v has nothing to give %v.\n", c.name, to.name)
		return
	}

	if to.leftHand != nil {
		fmt.Printf("%v's hands are full.\n", to.name)
		return
	}

	to.leftHand = c.leftHand
	fmt.Printf("%v gives a %v to %v.\n", c.name, c.leftHand.name, to.name)
	c.leftHand = nil
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v is carrying nothing.", c.name)
	}
	return fmt.Sprintf("%v is carrying a %v.", c.name, c.leftHand.name)
}