package gnresolve

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

// outputError outputs errors arrived from the name-finding process.
func (r GNresolve) outputError(errCh <-chan error, wgOut *sync.WaitGroup) {
	f, err := os.Create(filepath.Join(r.OutputPath, "errors.csv"))
	defer wgOut.Done()
	if err != nil {
		log.Fatal(err)
	}
	ef := csv.NewWriter(f)
	defer f.Close()
	defer ef.Flush()

	ef.Write([]string{"TimeStamp", "Error"})
	for e := range errCh {
		ef.Write([]string{ts(), e.Error()})
	}
}

// outputResults outputs data about found names.
func (r GNresolve) outputResult(outCh <-chan Verification,
	wgOut *sync.WaitGroup) {
	defer wgOut.Done()
	count := 0
	startTime := time.Now()

	f, err := os.Create(filepath.Join(r.OutputPath, "verification.csv"))
	if err != nil {
		log.Fatal(err)
	}

	of := csv.NewWriter(f)
	of.Write([]string{
		"Name", "DataSource", "DataSourceID", "Match", "MatchedName", "CurrentName",
		"Path",
	})

	defer f.Close()
	defer of.Flush()

	for v := range outCh {
		count++
		if r.ProgressNum > 0 && count%r.ProgressNum == 0 {
			rate := float64(count) / (time.Since(startTime).Minutes())
			log.Printf("Processing %dth name. Rate %0.2f names/min\n", count, rate)
		}
		out := []string{
			v.name, v.dataSource, v.dataSourceID, v.match, v.matchName,
			v.currentName, v.path,
		}
		of.Write(out)

		if err := of.Error(); err != nil {
			log.Fatal(err)
		}
	}
}

// ts generates a converted to a string timestamp in nanoseconds from epoch.
func ts() string {
	t := time.Now()
	return strconv.FormatInt(t.UnixNano(), 10)
}
