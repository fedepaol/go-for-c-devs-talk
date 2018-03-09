type Dog struct {
	Animal
}

func (d Dog) bark() {
	// bark
}

func (d *Dog) feed() {
	d.weight++
}

