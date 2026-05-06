package types

import "fmt"

/*
Build a report generator:
ReportGenerator interface with FetchData() string, FormatData(data string) string, Export(data string)
Run(r ReportGenerator) — the template method, always calls Fetch → Format → Export
Two concrete types — PDFReport and ExcelReport, each with different implementations of all three methods
In main — run both and show the different outputs
*/

type ReportGenerator interface {
	FetchData() string
	FormatData(string) string
	Export(string)
}

func Run(r ReportGenerator) {
	data := r.FetchData()
	formatted := r.FormatData(data)
	r.Export(formatted)
}

//PDFReport and ExcelReport
type PDFReport struct{}

func (p PDFReport) FetchData() string {
	return " pdf data fetched...."
}

func (p PDFReport) FormatData(data string) string {
	return " pdf data formatted...."
}

func (p PDFReport) Export(data string) {
	fmt.Println(" pdf data exported....")
}

type ExcelReport struct{}

func (p ExcelReport) FetchData() string {
	return " excel data fetched...."
}

func (p ExcelReport) FormatData(data string) string {
	return " excel data formatted...."
}

func (p ExcelReport) Export(data string) {
	fmt.Println(" excel data exported....")
}
