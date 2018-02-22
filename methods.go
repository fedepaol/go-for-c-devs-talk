type Dog struct {
	Animal
}

func (d Dog) bark() {
	// bark
}

func (d *Dog) feed() {
	d.weight++
}

var d Dog
d.bark()
d.feed() // (*d).feed() 
