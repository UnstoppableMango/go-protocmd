package cli_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	cliv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/unmango/cli/v1alpha1"
	"github.com/unstoppablemango/go-protocmd/pkg/cli"
)

var _ = Describe("Parse", func() {
	It("sets Raw to the original string", func() {
		result, err := cli.Parse("--verbose")

		Expect(err).NotTo(HaveOccurred())
		Expect(result.GetRaw()).To(Equal("--verbose"))
	})

	It("returns empty tokens for empty input", func() {
		result, err := cli.Parse("")

		Expect(err).NotTo(HaveOccurred())
		Expect(result.GetTokens()).To(BeEmpty())
	})

	It("returns empty tokens for whitespace-only input", func() {
		result, err := cli.Parse("   ")

		Expect(err).NotTo(HaveOccurred())
		Expect(result.GetTokens()).To(BeEmpty())
	})

	Describe("long flags", func() {
		It("parses a single long flag", func() {
			result, err := cli.Parse("--verbose")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("verbose"))
		})

		It("parses multiple long flags in order", func() {
			result, err := cli.Parse("--foo --bar")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(2))
			Expect(result.GetTokens()[0].GetLongFlag().GetName()).To(Equal("foo"))
			Expect(result.GetTokens()[1].GetLongFlag().GetName()).To(Equal("bar"))
		})

		It("parses long flag with inline value assignment", func() {
			result, err := cli.Parse("--output=file.txt")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(lf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_INLINE))
		})

		It("parses long flag with single-quoted inline value", func() {
			result, err := cli.Parse("--output='file.txt'")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(lf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_INLINE))
			Expect(lf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_SINGLE))
		})

		It("parses long flag with double-quoted inline value", func() {
			result, err := cli.Parse(`--output="file.txt"`)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(lf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_INLINE))
			Expect(lf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_DOUBLE))
		})
	})

	Describe("short flags", func() {
		It("parses a single short flag", func() {
			result, err := cli.Parse("-v")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"v"}))
		})

		It("parses a combined short flag cluster as one token", func() {
			result, err := cli.Parse("-abc")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"a", "b", "c"}))
		})
	})

	Describe("words", func() {
		It("parses a bare word", func() {
			result, err := cli.Parse("foo")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			w := result.GetTokens()[0].GetWord()
			Expect(w).NotTo(BeNil())
			Expect(w.GetValue()).To(Equal("foo"))
		})

		It("parses a single-quoted word", func() {
			result, err := cli.Parse("'foo'")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			w := result.GetTokens()[0].GetWord()
			Expect(w).NotTo(BeNil())
			Expect(w.GetValue()).To(Equal("foo"))
			Expect(w.GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_SINGLE))
		})

		It("parses a double-quoted word", func() {
			result, err := cli.Parse(`"foo"`)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			w := result.GetTokens()[0].GetWord()
			Expect(w).NotTo(BeNil())
			Expect(w.GetValue()).To(Equal("foo"))
			Expect(w.GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_DOUBLE))
		})
	})

	It("parses a long flag followed by a word", func() {
		result, err := cli.Parse("--verbose input.txt")

		Expect(err).NotTo(HaveOccurred())
		Expect(result.GetTokens()).To(HaveLen(2))
		Expect(result.GetTokens()[0].GetLongFlag()).NotTo(BeNil())
		Expect(result.GetTokens()[0].GetLongFlag().GetName()).To(Equal("verbose"))
		Expect(result.GetTokens()[1].GetWord()).NotTo(BeNil())
		Expect(result.GetTokens()[1].GetWord().GetValue()).To(Equal("input.txt"))
	})

	It("parses word, long flag, and word in order", func() {
		result, err := cli.Parse("cmd --flag arg")

		Expect(err).NotTo(HaveOccurred())
		Expect(result.GetTokens()).To(HaveLen(3))
		Expect(result.GetTokens()[0].GetWord()).NotTo(BeNil())
		Expect(result.GetTokens()[0].GetWord().GetValue()).To(Equal("cmd"))
		Expect(result.GetTokens()[1].GetLongFlag()).NotTo(BeNil())
		Expect(result.GetTokens()[1].GetLongFlag().GetName()).To(Equal("flag"))
		Expect(result.GetTokens()[2].GetWord()).NotTo(BeNil())
		Expect(result.GetTokens()[2].GetWord().GetValue()).To(Equal("arg"))
	})
})
