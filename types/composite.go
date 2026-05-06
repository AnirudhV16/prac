package types

import "fmt"

/*
Build an organization hierarchy:

Employee interface with GetName() string, GetSalary() float64, Display(indent string)
Developer struct — leaf node, has name and salary
Manager struct — composite node, has name, salary, and []Employee subordinates, plus AddSubordinate(e Employee)
Manager.GetSalary() returns their own salary plus the total salary of all subordinates
In main build a tree — one manager with two developers, that manager reports to a senior manager who also has one direct developer
*/

type Employee interface {
	GetName() string
	GetSalary() float64
	Display(string)
}

type Developer struct {
	name   string
	salary float64
}

func NewDeveloper(name string, salary float64) *Developer {
	return &Developer{
		name:   name,
		salary: salary,
	}
}

func (d *Developer) GetName() string {
	return d.name
}

func (d *Developer) GetSalary() float64 {
	return d.salary
}

func (d *Developer) Display(intent string) {
	fmt.Println(intent)
}

type Manager struct {
	name         string
	salary       float64
	subordinates []Employee
}

func NewManager(name string, salary float64, subordinates []Employee) *Manager {
	return &Manager{
		name:         name,
		salary:       salary,
		subordinates: subordinates,
	}
}

func (m *Manager) AddSubordinate(e Employee) {
	m.subordinates = append(m.subordinates, e)
}

func (m *Manager) GetName() string {
	return m.name
}

func (m *Manager) GetSalary() float64 {
	sum := m.salary
	for _, v := range m.subordinates {
		sum += v.GetSalary()
	}
	return sum
}

func (m *Manager) Display(intent string) {
	fmt.Println(intent)
}
