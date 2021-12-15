package main

type SortByArtist []*Track

func (x SortByArtist) Len() int {
	return len(x)
}

func (x SortByArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x SortByArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type CustomSort struct {
	t []*Track
	less func(x,y *Track) bool
}

func (x CustomSort) Len() int {
	return len(x.t)
}

func (x CustomSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}

func (x CustomSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}