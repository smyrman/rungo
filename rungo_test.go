package dispatcher

import (
	"testing"
	"time"
)

type Term bool

func (t Term) Run(termc <-chan bool) {
	<-termc
}

type SelectTerm bool

func (t SelectTerm) Run(termc <-chan bool) {
	tick := time.Tick(100*time.Millisecond)
	for {
		select {
		case <-termc:
			return
		case <-tick:
			// pass
		}
	}
}

type Wait bool

func (t Wait) Run(termc <-chan bool) {
	for i:=0; i < 99; i++ {
		// pass
	}
}

func TestTerm(t *testing.T) {
	var term1 Term
	var term2 SelectTerm
	routine1 := Go(term1)
	routine2 := Go(term2)
	routine1.Terminate()
	routine2.Terminate()
}

func TestTerm5x(t *testing.T) {
	var term Term
	routine := Go(term)
	routine.Terminate()
	// The rest of the calls should have no effect, and they shoud not result in a deadlock!
	routine.Terminate()
	routine.Terminate()
	routine.Terminate()
	routine.Terminate()
}

func TestWait(t *testing.T) {
	var wait Wait
	routine := Go(wait)
	routine.Wait()
}
