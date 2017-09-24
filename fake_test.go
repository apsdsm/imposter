package imposter_test

import (
	"github.com/apsdsm/imposter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Foo struct {
	imposter.Fake
}

type DoIt2Sig struct {
	A int
	B bool
	C string
}

func (f *Foo) DoIt() {
	f.SetCall("DoIt")
}

func (f *Foo) DoIt2(a int, b bool, c string) {
	f.SetCall("DoIt2", DoIt2Sig{a, b, c})
}

var _ = Describe("Fake", func() {

	var (
		f Foo
	)

	BeforeEach(func() {
		f = Foo{}
	})

	Context("Received", func() {
		It("return true if method with no parameters was called", func() {
			f.DoIt()

			Expect(f.Received("DoIt")).To(BeTrue())
		})

		It("returns true if method with parameters was called", func() {
			f.DoIt2(1, true, "foobar")

			Expect(f.Received("DoIt2", DoIt2Sig{1, true, "foobar"})).To(BeTrue())
		})

		It("panics if no methods were called", func() {
			Expect(func() {
				f.Received("DoIt")
			}).To(Panic())
		})

		It("panics if method was not called (but different method was called)", func() {
			f.DoIt2(1, true, "foo")

			Expect(func() {
				f.Received("DoIt")
			}).To(Panic())
		})

		It("panics if method with parameters was not called (but same method with different parameters was called)", func() {
			f.DoIt2(1, true, "foo")

			Expect(func() {
				f.Received("DoIt2", DoIt2Sig{2, false, "bar"})
			}).To(Panic())
		})
	})

	Context("DidNotReceive", func() {

		It("returns true if no methods were called", func() {
			Expect(f.DidNotReceive("DoIt")).To(BeTrue())
		})

		It("returns true if method not called (but different method was called)", func() {
			f.DoIt2(1, true, "foo")

			Expect(f.DidNotReceive("DoIt")).To(BeTrue())
		})

		It("returns true if method was called with different parameters", func() {
			f.DoIt2(1, true, "foo")

			Expect(f.DidNotReceive("DoIt2", DoIt2Sig{2, false, "bar"})).To(BeTrue())
		})

		It("panics if method was called", func() {
			f.DoIt()

			Expect(func() {
				f.DidNotReceive("DoIt")
			}).To(Panic())
		})

		It("panics if method with parameters was called", func() {
			f.DoIt2(1, true, "foo")

			Expect(func() {
				f.DidNotReceive("DoIt2", DoIt2Sig{1, true, "foo"})
			}).To(Panic())
		})
	})

})
