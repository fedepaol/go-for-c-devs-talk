c, ok := d.(Car) // true
c, ok := d.(Bike)  // false

switch d.(type) {
	case Car: //
}
