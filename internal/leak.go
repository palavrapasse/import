package internal

type Context string

type Leak struct {
	leakId      int
	ShareDateSC int
	Context     Context
}
