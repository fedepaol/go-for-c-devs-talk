type Driveable interface {
	func Drive() int
}

func (c Car) Drive (km int) {/* */}
func (b Bike) Drive (km int) {/* */}

var c Car{wheels : 4}
var d Driveable = c
