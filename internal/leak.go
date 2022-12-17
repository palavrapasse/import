package internal

type Context string

type Leak struct {
	LeakId      uint
	ShareDateSC int
	Context     Context
}
