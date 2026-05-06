package types

import "fmt"

/*
Build a home theatre facade:

Three subsystems — DVDPlayer with Play(movie string), Projector with TurnOn() and SetInput(input string), SoundSystem with SetVolume(level int) and TurnOn()
HomeTheatreFacade that composes all three
Two facade methods — WatchMovie(movie string) which turns everything on, sets input, sets volume to 8, and plays the movie. And EndMovie() which turns everything off
In main, just call WatchMovie and EndMovie
*/

type DVDPlayer struct{}

func (dvd *DVDPlayer) Play(movie string) {
	fmt.Printf("movie playing....\n")
}

type Projector struct{}

func (p *Projector) TurnOn() {
	fmt.Printf("projector turned on....\n")
}

func (p *Projector) SetInput(input string) {
	fmt.Printf("input okkk....\n")
}

type SoundSystem struct{}

func (s *SoundSystem) SetVolume(level int) {
	fmt.Printf("volume set to %d level\n", level)
}

func (s *SoundSystem) TurnOn() {
	fmt.Printf("speakers are turned on")
}

type HomeTheatreFacade struct {
	dvd       DVDPlayer
	projector Projector
	sound     SoundSystem
}

func (h *HomeTheatreFacade) WatchMovie(movie string) {
	h.dvd.Play(movie)
	h.projector.TurnOn()
	h.sound.TurnOn()
	h.projector.SetInput("USB")
	h.sound.SetVolume(8)
}

func (h *HomeTheatreFacade) EndMovie() {
	fmt.Printf("stop playing....\n")
}
