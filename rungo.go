// A small library that lets you dispatch goroutines, and then enables you to
// either wait until they finish or ask for their safe termination.
package rungo

import (
	"sync"
)

// Implementations of this interface's Run function should periodically check
// the termc channel and return if there is something on it.
type Interface interface {
	Run(termc <-chan bool)
}

// A dispatcher object that can be used to safly terminate a goroutine.
type Routine struct {
	wg sync.WaitGroup
	termc chan bool
	I Interface

}

// Call I.Run() in a new goroutine, and return a dispatcher object that can be
// used to wait for it's completion, or to ask for it's termination.
func Go(I Interface) *Routine {
	r := new(Routine)
	r.termc = make(chan bool, 1)
	r. I = I
	r.wg.Add(1)
	go r.run()
	return r
}

func (r *Routine) run() {
	defer r.wg.Done()
	r.I.Run(r.termc)
}

// If the termc channel of the Interface is not full, signal it. Block until
// the routine has ended.  This call should be considered thread-safe. If  the
// routine has already ended, or Terminate has been called before (possibly
// from another goroutine), this function returns at once.
func (r *Routine) Terminate() {
	select {
	default:
		// pass
	case r.termc <- false:
		r.wg.Wait()
	}
	return
}

// Block until the routine has ended.
func (r *Routine) Wait() {
	r.wg.Wait()
}
