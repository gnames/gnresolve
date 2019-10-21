package gnresolve

import (
	"strconv"
	"sync"

	"github.com/gnames/gnfinder/verifier"
)

func (r GNresolve) worker(inCh <-chan []string, outCh chan<- Verification,
	errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	var sources []int
	verif := verifier.NewVerifier(verifier.OptWorkers(r.JobsNum),
		verifier.OptSources(sources))

	for names := range inCh {
		verified := verif.Run(names)
		processVerified(verified, outCh)
	}
}

func processVerified(verified verifier.Output, outCh chan<- Verification) {
	for k, v := range verified {
		br := v.BestResult
		res := Verification{
			name:         k,
			dataSource:   br.DataSourceTitle,
			dataSourceID: strconv.Itoa(br.DataSourceID),
			match:        br.MatchType,
			editDistance: strconv.Itoa(br.EditDistance),
			matchName:    br.MatchedName,
			currentName:  br.CurrentName,
			path:         br.ClassificationPath,
			ranks:        br.ClassificationRank,
		}
		outCh <- res
	}
}
