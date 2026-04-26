package read

//go:generate stringer -type=tokenType
type tokenType int

const (
	illegal tokenType = iota
	rParen
	lParen
	quote
	number
	boolean
	symbol
	eof
)

type token struct {
	typ tokenType
	str string
}
