package conv_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unstoppablemango/go-protocmd/pkg/conv"
	"github.com/unstoppablemango/go-protocmd/pkg/conv/builder"
)

var _ = Describe("Builder", func() {
	It("should append strings", func() {
		b := builder.New()

		conv.Append(b, "test")

		Expect(b.Build()).To(HaveExactElements("test"))
	})

	It("should append integers", func() {
		b := builder.New()

		conv.Append(b, 69)

		Expect(b.Build()).To(HaveExactElements("69"))
	})

	It("should append values", func() {
		b := builder.New()

		conv.Append(b, 420, 69)

		Expect(b.Build()).To(HaveExactElements("420", "69"))
	})

	It("should append values in subsequent calls", func() {
		b := builder.New()

		conv.Append(b, 420)
		conv.Append(b, 69)

		Expect(b.Build()).To(HaveExactElements("420", "69"))
	})

	It("should conditionally append", func() {
		b := builder.New()

		conv.AppendIf(b, true, "69")

		Expect(b.Build()).To(HaveExactElements("69"))
	})

	It("should conditionally skip", func() {
		b := builder.New()

		conv.AppendIf(b, false, "69")

		Expect(b.Build()).To(BeEmpty())
	})
})
