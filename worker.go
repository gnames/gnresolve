package gnresolve

import "sync"

func (r GNresolve) worker(inCh <-chan string, outCh chan<- Verification,
	errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range inCh {
		res := Verification{name: n}
		outCh <- res
	}
}
