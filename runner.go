package gnresolve

import (
	"bufio"
	"os"
	"sync"
)

const BATCH_SIZE = 50000

type Verification struct {
	name         string
	dataSource   string
	dataSourceID string
	match        string
	editDistance string
	matchName    string
	currentName  string
	path         string
	ranks        string
}

func (r GNresolve) Run() error {
	inCh := make(chan []string)
	errCh := make(chan error)
	outCh := make(chan Verification)
	var wg sync.WaitGroup
	var wgOut sync.WaitGroup
	wg.Add(1)
	wgOut.Add(2)
	go r.outputError(errCh, &wgOut)
	go r.outputResult(outCh, &wgOut)
	for i := 0; i < 1; i++ {
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

func (r GNresolve) streamInput(inCh chan<- []string, errCh chan<- error) error {
	f, err := os.Open(r.InputPath)
	if err != nil {
		return err
	}
	defer f.Close()
	defer close(inCh)

	names := make([]string, 0, BATCH_SIZE)

	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if count > BATCH_SIZE {
			count = 0
			inCh <- names
			names = make([]string, 0, BATCH_SIZE)
		}
		names = append(names, scanner.Text())
		count++
	}
	if err = scanner.Err(); err != nil {
		errCh <- err
	}
	return nil
}
