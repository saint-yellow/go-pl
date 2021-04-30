package main

type radio struct {
	fromRover chan message
}

func (r *radio) sendToEarth(m message) {
	r.fromRover <- m
}

func newRadio(toEarth chan []message) *radio {
	r := &radio{
		fromRover: make(chan message),
	}
	go r.run(toEarth)
	return r
}

func (r *radio) run(toEarth chan []message) {
	var buffered []message
	for {
		toEarth1 := toEarth
		if len(buffered) == 0 {
			toEarth1 = nil
		}
		select {
		case m := <-r.fromRover:
			buffered = append(buffered, m)
		case toEarth1 <- buffered:
			buffered = nil
		}
	}
}