package main

func (p byName) Len() int {
	return len(p)
}

func (p byName) Less(i, j int) bool {
	return p[i].name < p[j].name
}

func (p byName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
