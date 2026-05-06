package use

import (
	"fmt"
	"time"
)

/*
Design a parking lot system that can park and unpark vehicles, track available spots, and generate a ticket when a vehicle parks.
Requirements

The parking lot has multiple floors, each floor has multiple spots
Spots are of three types — Bike, Car, Truck — each vehicle type can only park in its spot type
When a vehicle parks, generate a ticket with spot info and entry time
When a vehicle unparks, calculate the fee based on hours parked — Bike: ₹10/hr, Car: ₹20/hr, Truck: ₹40/hr
The system should find the nearest available spot for a given vehicle type
The parking lot should report how many spots are available per type
*/

/*
1. parking lot - floors []slice of floor, park  and unpark functinalities, call the multiple types functionalities and hide complexity here??
 -- park() - will find the nearest spot of needed vehicle type and generates a ticket with spot and vehicle details
 -- unpack() - fee generation based on vehicle type, spot is marked as available again
2. floor - spots slice of spot[], spot status and type of spot, park() and unpark() , it will have functionality to find the best spot in floor??
3. vehicle - factory that will return asked concrete type of vehicle, this is used in ticket generation
4. ticket - spot, vehicle, fee generator(strategy)
5. finding best spot in a parking lot???? (a method of parking lot)
*/

// spot types
type Spot interface {
	GetSpotType() string
	isAvailable() bool
	Park()   //this will set the availability to false
	Unpark() //this will set the availability to true
	Calculate(time.Time) float64
}

type BasicSpot struct {
	available     bool
	spotType      string
	feeCalculator FeeCalculator
}

func NewBasicSpot(FeeCalculator FeeCalculator, typee string) *BasicSpot {
	t := &BasicSpot{feeCalculator: FeeCalculator, spotType: typee}
	t.available = true
	return t
}

func (b *BasicSpot) Calculate(t time.Time) float64 {
	return b.feeCalculator.Calculate(t)
}

func (b *BasicSpot) GetSpotType() string {
	return b.spotType
}

func (b *BasicSpot) isAvailable() bool {
	return b.available
}

func (b *BasicSpot) Park() {
	b.available = false
}

func (b *BasicSpot) Unpark() {
	b.available = true
}

// fee calculation
type FeeCalculator interface {
	Calculate(time.Time) float64
}

type BikeFeeCalculator struct{}

func (b BikeFeeCalculator) Calculate(t time.Time) float64 {
	return time.Since(t).Hours() * 10 //10 per hour
}

type CarFeeCalculator struct{}

func (b CarFeeCalculator) Calculate(t time.Time) float64 {
	return time.Since(t).Hours() * 20 //20 per hour
}

type TruckFeeCalculator struct{}

func (b TruckFeeCalculator) Calculate(t time.Time) float64 {
	return time.Since(t).Hours() * 30 //30 per hour
}

// vehicle types
type Vehicle struct {
	model  string
	number string
	typee  string
}

func NewVehicle(model string, number string, typee string) *Vehicle {
	return &Vehicle{model: model, number: number, typee: typee}
}

func (b *Vehicle) GetModel() string {
	return b.model
}

func (b *Vehicle) GetNumber() string {
	return b.number
}

func (b *Vehicle) GetType() string {
	return b.typee
}

type ParkingLot struct {
	Floors []Floor
}

func (p *ParkingLot) AvailableSpots() map[string]int {
	counts := map[string]int{"bike": 0, "car": 0, "truck": 0}
	for _, floor := range p.Floors {
		for _, spot := range floor.GetSpots() {
			if spot.isAvailable() {
				counts[spot.GetSpotType()]++
			}
		}
	}
	return counts
}

func (p *ParkingLot) Park(vehicle Vehicle) (*Ticket, error) {
	//finding spot in each floor
	for _, v := range p.Floors {
		ticket := v.FindSpot(vehicle) // this will set the spot aailability, this will make the ticket and returns it.
		if ticket != nil {
			ticket.Display() //displaying the ticket....
			return ticket, nil
		}
	}
	return nil, fmt.Errorf("parking lot is full....")
}

func (p *ParkingLot) Unpark(ticket Ticket) float64 {
	//change the availability of spot, then?? calculate the fee for this vehicle type from entry time
	//display that fee
	ticket.GetSpot().Unpark()
	return ticket.GetSpot().Calculate(ticket.GetEntryTime())
}

type Floor struct {
	spots []Spot
}

func (f *Floor) AddSpots(spot Spot) {
	f.spots = append(f.spots, spot)
}

func (f *Floor) GetSpots() []Spot {
	return f.spots
}

func (f *Floor) FindSpot(vehicle Vehicle) *Ticket {
	//go through all floors, and all spots get the available spot of needed vehicle type
	// how to check the vehicle type??
	for _, v := range f.spots {
		if v.GetSpotType() == vehicle.typee {
			if v.isAvailable() == true {
				// this should generate ticket and return that???
				ticket := NewTicket(vehicle, time.Now().Add(-2*time.Hour), v)
				v.Park()
				return ticket
			}
		}
	}
	return nil
}

type Ticket struct {
	vehicle   Vehicle
	spot      Spot
	entryTime time.Time
}

func (t *Ticket) GetVehicle() Vehicle     { return t.vehicle }
func (t *Ticket) GetSpot() Spot           { return t.spot }
func (t *Ticket) GetEntryTime() time.Time { return t.entryTime }

func NewTicket(vehicle Vehicle, t time.Time, spot Spot) *Ticket {
	return &Ticket{vehicle: vehicle, entryTime: t, spot: spot}
}

func (t *Ticket) Display() {
	fmt.Printf("ticket generated, entry time %d\n", t.entryTime.Hour())
}
