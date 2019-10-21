package gnresolve

import (
	"fmt"
	"os"
)

type GNresolve struct {
	// InputPath gives path to a file with an input data. The file should
	// contain name-strings separated by new lines.
	InputPath string
	// OutputPath gives path to a directory to keep output data.
	OutputPath string
	// JobsNum sets number of jobs/workers to run.
	JobsNum int
	// ProgressNum determines how many titles should be processed for
	// a progress report.
	ProgressNum int
}

type Option func(r *GNresolve)

// OptJobs sets number of jobs/workers to run duing execution.
func OptJobs(i int) Option {
	return func(r *GNresolve) {
		r.JobsNum = i
	}
}

// OptProgressNum sets how often to printout a line about the progress. When it
// is set to 1 report line appears after processing every title, and if it is 10
// progress is shows after every 10th title.
func OptProgressNum(i int) Option {
	return func(r *GNresolve) {
		r.ProgressNum = i
	}
}

// OptIntput is an absolute path to input data file. Each line of such file
// displays path to zipped file of a title.
func OptInput(s string) Option {
	return func(r *GNresolve) {
		r.InputPath = s
	}
}

// OptOutput is an absolute path to a directory where results will be written.
// If such directory does not exist already, it will be created during
// initialization of HTindex instance.
func OptOutput(s string) Option {
	return func(r *GNresolve) {
		r.OutputPath = s
	}
}

func NewGNresolve(opts ...Option) (GNresolve, error) {
	res := GNresolve{
		JobsNum:     5,
		ProgressNum: 10000,
		OutputPath:  "/tmp/gnresolve",
	}
	for _, opt := range opts {
		opt(&res)
	}
	err := res.setOutputDir()
	return res, err
}

func (r GNresolve) setOutputDir() error {
	path, err := os.Stat(r.OutputPath)
	if os.IsNotExist(err) {
		return os.MkdirAll(r.OutputPath, 0755)
	}
	if path.Mode().IsRegular() {
		return fmt.Errorf("'%s' is a file, not a directory", r.OutputPath)
	}
	return nil
}
