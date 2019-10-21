package gnresolve_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gnames/gnresolve"
)

var _ = Describe("GNresolve", func() {
	Describe("NewGNresolve", func() {
		It("creates new GNresolve with default settings", func() {
			r, err := NewGNresolve()
			Expect(err).To(BeNil())
			Expect(r.ProgressNum).To(Equal(10000))
		})

		It("creates new GNresolve using opts", func() {
			opts := []Option{
				OptInput("testdata/names.txt"),
				OptOutput("/tmp/gnroutput"),
				OptJobs(3),
				OptProgressNum(100),
			}
			r, err := NewGNresolve(opts...)
			Expect(err).To(BeNil())
			Expect(r.ProgressNum).To(Equal(100))
			Expect(r.OutputPath).To(Equal("/tmp/gnroutput"))
			Expect(r.JobsNum).To(Equal(3))
		})
	})

	Describe("#Run", func() {
		It("processes input", func() {
			r, err := makeGNresolve()
			Expect(err).To(BeNil())
			r.Run()
		})
	})
})

func makeGNresolve() (GNresolve, error) {
	opts := []Option{
		OptInput("testdata/names.txt"),
		OptOutput("/tmp/gnroutput"),
		OptJobs(3),
		OptProgressNum(1000),
	}
	return NewGNresolve(opts...)
}
