package ccy

type Ccy struct {
	sem chan bool
}

func (ccy *Ccy) Lock() {
	ccy.sem <- true
}

func (ccy *Ccy) Unlock() {
	<-ccy.sem
}
