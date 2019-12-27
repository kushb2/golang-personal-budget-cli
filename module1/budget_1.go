package module1

// Budget stores budget information
type budget struct {
	max   float32
	items []item
}

// Item stores item information
type item struct {
	description string
	price       float32
}
