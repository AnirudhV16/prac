package types

import "fmt"

//location part that will take location coordinates and a function returns them as a string
type Location struct {
	lat float64
	lng float64
}

func (l *Location) Coordinates() string {
	return fmt.Sprintf("%.4f , %.4f", l.lat, l.lng)
}

//Rating struct
type Rating struct {
	score        float64
	totalReviews int
}

func (r *Rating) AddReview(score float64) {
	r.totalReviews++
	r.score = (r.score*(float64(r.totalReviews-1)) + score) / float64(r.totalReviews)
}

func (r *Rating) GetRating() float64 {
	return r.score
}

//restaurent struct with locationa and rating structs
type Restaurant struct {
	name string
	Location
	Rating
}

func NewRestaurtant(name string, lat float64, lng float64, score float64, totalReviews int) *Restaurant {
	return &Restaurant{
		name:     name,
		Location: Location{lat: lat, lng: lng},
		Rating:   Rating{score: score, totalReviews: totalReviews}}
}

//Ratable interface that prints the rating of any item
type Rateable interface {
	AddReview(score float64)
	GetRating() float64
}

func PrintTopRated(items []Rateable, threshold float64) {
	for _, v := range items {
		if v.GetRating() > threshold {
			fmt.Println(v.GetRating())
		}
	}
}
