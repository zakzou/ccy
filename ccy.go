package ccy

type Ccy struct {
	sem chan bool
}

func NewCcy(sem chan bool) *Ccy {
	ccy := new(Ccy)
	ccy.sem = sem
	return ccy
}

func (ccy *Ccy) Lock() {
	ccy.sem <- true
}

func (ccy *Ccy) Unlock() {
	<-ccy.sem
}
