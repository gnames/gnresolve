package gnresolve

import (
	"bufio"
	"os"
	"sync"
)

type Verification struct {
	name         string
	dataSource   string
	dataSourceID string
	match        string
	matchName    string
	currentName  string
	path         string
}

func (r GNresolve) Run() error {
	inCh := make(chan string)
	errCh := make(chan error)
	outCh := make(chan Verification)
	var wg sync.WaitGroup
	var wgOut sync.WaitGroup
	wg.Add(r.JobsNum)
	wgOut.Add(2)
	go r.outputError(errCh, &wgOut)
	go r.outputResult(outCh, &wgOut)
	for i := 0; i < r.JobsNum; i++ {
		go r.worker(inCh, outCh, errCh, &wg)
	}
	if err := r.streamInput(inCh, errCh); err != nil {
		return err
	}
	wg.Wait()
	close(outCh)
	close(errCh)
	wgOut.Wait()
	return nil
}

func (r GNresolve) streamInput(inCh chan<- string, errCh chan<- error) error {
	f, err := os.Open(r.InputPath)
	if err != nil {
		return err
	}
	defer f.Close()
	defer close(inCh)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		inCh <- scanner.Text()
	}
	if err = scanner.Err(); err != nil {
		errCh <- err
	}
	return nil
}
