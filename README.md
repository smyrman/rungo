Rungo
=====

Rungo is a small library that lets you dispatch goroutines, and then enables
you to either wait until they finish or ask for their safe termination.


Installation
------------

Issue:

	$ go get github.com/smyrman/rungo


Usage
-----

Define a rungo.Interface compatible type:

	type T bool
	func (t T) Run(termc <-chan bool) {
		var i int64
		defer fmt.Println("exit, i =", i)
		for i=0; i< 1e9; i++ {
			select {
			case <-termc:
				return
			default:
			}
			i++
		}
		return
	}

Exmaple uage of Terminate():

	var test T
	r := rungo.Go(test)
	defer r.Terminate()
	// <- Do some stuff ...

Exmaple uage of Wait():

	var test T
	r := rungo.Go(test)
	defer r.Wait()
	// <- Do some stuff ...

If you need more ways to communicate with the running goroutine, you could
store channels within your rungo.Interface compatible type:

	type T struct {
		C1 chan int
		C2 chan int
	}
	func NewT() *T {...}
	func (t T) Run(termc <-chan bool) { ... }


Origin
------

The rungo package was originally developed for use in the robot Loke
(eng:Loki), that participated in the Eurobot competition in France in 2012
(http://eurobot-ntnu.no).
