package internal

type Context string

type Leak struct {
	LeakId      AutoGenKey
	ShareDateSC int
	Context     Context
}
