package types

import "fmt"

type Printer interface {
	Print(string)
}

type Scanner interface {
	Scan(string)
}

type FaxMachine interface {
	Fax(string)
}

type DuplexPrinter interface {
	PrintDuplex(string)
}

// basic printer -> only print
type BasicPrinter struct {
	name string
}

func (b *BasicPrinter) Print(s string) {
	fmt.Printf("%s is printing", b.name)
}

//office printer -> print,scan,fac,duplex print

type FullPrinter interface {
	Printer
	Scanner
	FaxMachine
	DuplexPrinter
}

type OfficePrinter struct {
	name string
}

func (o *OfficePrinter) Print(s string) {
	fmt.Printf("%s is printing", o.name)
}

func (o *OfficePrinter) Scan(s string) {
	fmt.Printf("%s is scanning", o.name)
}

func (o *OfficePrinter) Fax(s string) {
	fmt.Printf("%s is Faxing", o.name)
}

func (o *OfficePrinter) PrintDuplex(s string) {
	fmt.Printf("%s is Printing Duplex", o.name)
}
