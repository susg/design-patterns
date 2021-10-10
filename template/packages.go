package template

import "fmt"

// malaysiaPackage is `trip` subclass which implements `iTrip`
type malaysiaPackage struct {
	trip
	source string
}

func newMalaysiaPackage(s string) malaysiaPackage {
	t := trip{}
	mp := malaysiaPackage{t, s}
	mp.iTrip = mp
	return mp
}

func (p malaysiaPackage) day1Activities() string {
	return "Malaysia day1"
}

func (p malaysiaPackage) day2Activities() string {
	return "Malaysia day2"
}

func (p malaysiaPackage) returnHome() string {
	return fmt.Sprintf("return %s", p.source)
}

func (p malaysiaPackage) setVehicleType() string {
	return "need SUV"
}

func (p malaysiaPackage) isRestNeeded() bool {
	return false
}

// dubaiPackage is `trip` subclass which implements `iTrip`
type dubaiPackage struct {
	trip
	source string
}

func newDubaiPackage(s string) dubaiPackage {
	t := trip{}
	dp := dubaiPackage{t, s}
	dp.iTrip = dp
	return dp
}

func (p dubaiPackage) day1Activities() string {
	return "Dubai day1"
}

func (p dubaiPackage) day2Activities() string {
	return "Dubai day2"
}

func (p dubaiPackage) returnHome() string {
	return fmt.Sprintf("return %s", p.source)
}

func (p dubaiPackage) setVehicleType() string {
	return "need Sedan"
}

func (p dubaiPackage) isRestNeeded() bool {
	return true
}
