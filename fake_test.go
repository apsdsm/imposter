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
	It("reports if method with no parameters was called", func() {
		f := Foo{}

		f.DoIt()

		Expect(f.Received("DoIt")).To(BeTrue())
	})

	It("reports if method with parameters was called", func() {
		f := Foo{}

		f.DoIt2(1, true, "foobar")

		Expect(f.Received("DoIt2", DoIt2Sig{1, true, "foobar"})).To(BeTrue())
	})
})
