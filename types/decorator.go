package types

import (
	"strings"
)

/*
Build a text processor system:

TextProcessor interface with Process(text string) string
PlainTextProcessor — base implementation, just returns the text as-is
UpperCaseDecorator — wraps a TextProcessor, converts output to uppercase
TrimDecorator — wraps a TextProcessor, trims whitespace from output
ExclamationDecorator — wraps a TextProcessor, adds !!! to the end
In main, chain all three decorators and process the text "  hello world  "

Expected output: HELLO WORLD!!!
*/

type TextProcessor interface {
	Process(string) string
}

type PlainTextProcessor struct{}

func NewPlainTextProcessor() *PlainTextProcessor {
	return &PlainTextProcessor{}
}

func (p PlainTextProcessor) Process(s string) string {
	return s
}

type UpperCaseDecorator struct {
	plain TextProcessor
}

func NewUpperCaseDecorator(plain TextProcessor) *UpperCaseDecorator {
	return &UpperCaseDecorator{plain: plain}
}

func (p UpperCaseDecorator) Process(s string) string {
	v := p.plain.Process(s)
	return strings.ToUpper(v)
}

type TrimDecorator struct {
	upper TextProcessor
}

func NewTrimDecorator(upper TextProcessor) *TrimDecorator {
	return &TrimDecorator{upper: upper}
}

func (p TrimDecorator) Process(s string) string {
	v := p.upper.Process(s)
	return strings.Trim(v, " ")
}

type ExclamationDecorator struct {
	trim TextProcessor
}

func NewExclamationDecorator(trim TextProcessor) *ExclamationDecorator {
	return &ExclamationDecorator{trim: trim}
}

func (p ExclamationDecorator) Process(s string) string {
	v := p.trim.Process(s)
	return v + "!!!"
}
