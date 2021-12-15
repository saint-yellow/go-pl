package main

import "fmt"

type People struct {
	Name string
	age int
}

func (p *People) SetName(name string) {
	p.Name = name
}

func (p *People) SetAge(age int) {
	p.age = age
}

func (p *People) GetAge() int {
	return p.age
}

type Kid struct {
	Name string
	Age int
}

func (k *Kid) SetName(name string) {
	k.Name = name
}

func (k *Kid) SetAge(age int) {
	k.Age = age
}

func main() {
	p := People{Name: "saint-yellow", age: 28}
	
	fmt.Printf("name: %s, age: %d\n", p.Name, p.age)

	p.SetName("Saint-Yellow Huang")
	p.SetAge(23)
	fmt.Printf("name: %s, age: %d\n", p.Name, p.age)
}