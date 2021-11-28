package template

import "fmt"

// malaysiaPackage is `trip` subclass which implements `iTrip`
type malaysiaPackage struct {
	trip
	source     string
	restNeeded bool
}

func newMalaysiaPackage(s string, rn bool) malaysiaPackage {
	t := trip{}
	mp := malaysiaPackage{t, s, rn}
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
	return p.restNeeded
}

// dubaiPackage is `trip` subclass which implements `iTrip`
type dubaiPackage struct {
	trip
	source     string
	restNeeded bool
}

func newDubaiPackage(s string, rn bool) dubaiPackage {
	t := trip{}
	dp := dubaiPackage{t, s, rn}
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
	return p.restNeeded
}
