type Vehchle struct {
	wheels int
}

type Bike struct {
	Vehicle
	pedals int
}

var b Bike
b.wheels // b.Vehicle.wheels
