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
		Expect(result.GetRaw()).To(Equal(""))
	})

	It("returns empty tokens for whitespace-only input", func() {
		result, err := cli.Parse("   ")

		Expect(err).NotTo(HaveOccurred())
		Expect(result.GetTokens()).To(BeEmpty())
		Expect(result.GetRaw()).To(Equal("   "))
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
			Expect(lf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_NONE))
		})

		It("parses long flag with space-separated value", func() {
			result, err := cli.Parse("--output file.txt")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(lf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_SPACE))
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

		It("parses long flag with space-separated single-quoted value", func() {
			result, err := cli.Parse("--output 'file.txt'")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(lf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_SPACE))
			Expect(lf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_SINGLE))
		})

		It("parses long flag with space-separated double-quoted value", func() {
			result, err := cli.Parse(`--output "file.txt"`)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(lf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_SPACE))
			Expect(lf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_DOUBLE))
		})

		It("parses long flag with hyphenated name", func() {
			result, err := cli.Parse("--no-cache")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("no-cache"))
		})

		It("parses long flag with hyphenated name and inline value", func() {
			result, err := cli.Parse("--dry-run=true")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("dry-run"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal("true"))
			Expect(lf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_INLINE))
		})

		It("parses long flag with empty inline value", func() {
			result, err := cli.Parse("--output=")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal(""))
			Expect(lf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_INLINE))
			Expect(lf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_NONE))
		})

		It("parses long flag with inline double-quoted value containing spaces", func() {
			result, err := cli.Parse(`--output="file name.txt"`)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal("file name.txt"))
			Expect(lf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_INLINE))
			Expect(lf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_DOUBLE))
		})

		It("parses long flag with inline single-quoted value containing spaces", func() {
			result, err := cli.Parse("--output='file name.txt'")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal("file name.txt"))
			Expect(lf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_INLINE))
			Expect(lf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_SINGLE))
		})

		It("parses single-character long flag name", func() {
			result, err := cli.Parse("--v")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("v"))
		})

		It("parses long flag name with digits", func() {
			result, err := cli.Parse("--output2")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output2"))
		})

		It("parses long flag with space-separated single-quoted value containing spaces", func() {
			result, err := cli.Parse("--output 'file name.txt'")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal("file name.txt"))
			Expect(lf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_SPACE))
			Expect(lf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_SINGLE))
		})

		It("parses long flag with space-separated double-quoted value containing spaces", func() {
			result, err := cli.Parse(`--output "file name.txt"`)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal("file name.txt"))
			Expect(lf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_SPACE))
			Expect(lf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_DOUBLE))
		})

		It("parses long flag with empty inline single-quoted value", func() {
			result, err := cli.Parse("--output=''")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal(""))
			Expect(lf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_INLINE))
			Expect(lf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_SINGLE))
		})

		It("parses long flag with empty inline double-quoted value", func() {
			result, err := cli.Parse(`--output=""`)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			lf := result.GetTokens()[0].GetLongFlag()
			Expect(lf).NotTo(BeNil())
			Expect(lf.GetName()).To(Equal("output"))
			Expect(lf.GetAssignment()).NotTo(BeNil())
			Expect(lf.GetAssignment().GetValue()).To(Equal(""))
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

		It("parses short flag with inline value assignment", func() {
			result, err := cli.Parse("-o=file.txt")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_INLINE))
		})

		It("parses short flag with adjacent value (POSIX)", func() {
			result, err := cli.Parse("-ofile.txt")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_ADJACENT))
		})

		It("parses short flag with space-separated value", func() {
			result, err := cli.Parse("-o file.txt")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_SPACE))
		})

		It("parses short flag with single-quoted inline value", func() {
			result, err := cli.Parse("-o='file.txt'")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_INLINE))
			Expect(sf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_SINGLE))
		})

		It("parses short flag with double-quoted inline value", func() {
			result, err := cli.Parse(`-o="file.txt"`)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_INLINE))
			Expect(sf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_DOUBLE))
		})

		It("parses short flag with space-separated single-quoted value", func() {
			result, err := cli.Parse("-o 'file.txt'")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_SPACE))
			Expect(sf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_SINGLE))
		})

		It("parses short flag with space-separated double-quoted value", func() {
			result, err := cli.Parse(`-o "file.txt"`)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_SPACE))
			Expect(sf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_DOUBLE))
		})

		It("parses short flag with adjacent single-quoted value", func() {
			result, err := cli.Parse("-o'file.txt'")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_ADJACENT))
			Expect(sf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_SINGLE))
		})

		It("parses short flag with adjacent double-quoted value", func() {
			result, err := cli.Parse(`-o"file.txt"`)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("file.txt"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_ADJACENT))
			Expect(sf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_DOUBLE))
		})

		It("parses short flag with empty inline value", func() {
			result, err := cli.Parse("-o=")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal(""))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_INLINE))
			Expect(sf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_NONE))
		})

		It("parses short flag cluster with inline value on last flag", func() {
			result, err := cli.Parse("-abc=val")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"a", "b", "c"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("val"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_INLINE))
		})

		It("treats all-alpha characters after dash as flag name cluster with no assignment", func() {
			result, err := cli.Parse("-abcval")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"a", "b", "c", "v", "a", "l"}))
			Expect(sf.GetAssignment()).To(BeNil())
		})

		It("parses multiple separate short flags as separate tokens", func() {
			result, err := cli.Parse("-v -x")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(2))
			Expect(result.GetTokens()[0].GetShortFlag()).NotTo(BeNil())
			Expect(result.GetTokens()[0].GetShortFlag().GetNames()).To(Equal([]string{"v"}))
			Expect(result.GetTokens()[1].GetShortFlag()).NotTo(BeNil())
			Expect(result.GetTokens()[1].GetShortFlag().GetNames()).To(Equal([]string{"x"}))
		})

		It("parses short flag with space-separated single-quoted value containing spaces", func() {
			result, err := cli.Parse("-o 'file name.txt'")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("file name.txt"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_SPACE))
			Expect(sf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_SINGLE))
		})

		It("parses short flag with space-separated double-quoted value containing spaces", func() {
			result, err := cli.Parse(`-o "file name.txt"`)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("file name.txt"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_SPACE))
			Expect(sf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_DOUBLE))
		})

		It("parses short flag with adjacent single-quoted value containing spaces", func() {
			result, err := cli.Parse("-o'file name.txt'")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("file name.txt"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_ADJACENT))
			Expect(sf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_SINGLE))
		})

		It("parses short flag with adjacent double-quoted value containing spaces", func() {
			result, err := cli.Parse(`-o"file name.txt"`)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal("file name.txt"))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_ADJACENT))
			Expect(sf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_DOUBLE))
		})

		It("parses short flag with empty adjacent single-quoted value", func() {
			result, err := cli.Parse("-o''")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal(""))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_ADJACENT))
			Expect(sf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_SINGLE))
		})

		It("parses short flag with empty adjacent double-quoted value", func() {
			result, err := cli.Parse(`-o""`)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			sf := result.GetTokens()[0].GetShortFlag()
			Expect(sf).NotTo(BeNil())
			Expect(sf.GetNames()).To(Equal([]string{"o"}))
			Expect(sf.GetAssignment()).NotTo(BeNil())
			Expect(sf.GetAssignment().GetValue()).To(Equal(""))
			Expect(sf.GetAssignment().GetStyle()).To(Equal(cliv1alpha1.AssignmentStyle_ASSIGNMENT_STYLE_ADJACENT))
			Expect(sf.GetAssignment().GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_DOUBLE))
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
			Expect(w.GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_NONE))
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

		It("parses an empty single-quoted word", func() {
			result, err := cli.Parse("''")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			w := result.GetTokens()[0].GetWord()
			Expect(w).NotTo(BeNil())
			Expect(w.GetValue()).To(Equal(""))
			Expect(w.GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_SINGLE))
		})

		It("parses an empty double-quoted word", func() {
			result, err := cli.Parse(`""`)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			w := result.GetTokens()[0].GetWord()
			Expect(w).NotTo(BeNil())
			Expect(w.GetValue()).To(Equal(""))
			Expect(w.GetQuoteStyle()).To(Equal(cliv1alpha1.QuoteStyle_QUOTE_STYLE_DOUBLE))
		})

		It("parses multiple bare words as separate tokens", func() {
			result, err := cli.Parse("foo bar baz")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(3))
			Expect(result.GetTokens()[0].GetWord().GetValue()).To(Equal("foo"))
			Expect(result.GetTokens()[1].GetWord().GetValue()).To(Equal("bar"))
			Expect(result.GetTokens()[2].GetWord().GetValue()).To(Equal("baz"))
		})
	})

	Describe("end of options", func() {
		It("parses -- as an end-of-options token", func() {
			result, err := cli.Parse("--")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(1))
			Expect(result.GetTokens()[0].GetEndOfOptions()).NotTo(BeNil())
		})

		It("parses -- followed by a word as end-of-options then word", func() {
			result, err := cli.Parse("-- foo")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(2))
			Expect(result.GetTokens()[0].GetEndOfOptions()).NotTo(BeNil())
			Expect(result.GetTokens()[1].GetWord()).NotTo(BeNil())
			Expect(result.GetTokens()[1].GetWord().GetValue()).To(Equal("foo"))
		})

		It("parses flag then -- then word in order", func() {
			result, err := cli.Parse("--verbose -- positional")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(3))
			Expect(result.GetTokens()[0].GetLongFlag()).NotTo(BeNil())
			Expect(result.GetTokens()[0].GetLongFlag().GetName()).To(Equal("verbose"))
			Expect(result.GetTokens()[1].GetEndOfOptions()).NotTo(BeNil())
			Expect(result.GetTokens()[2].GetWord()).NotTo(BeNil())
			Expect(result.GetTokens()[2].GetWord().GetValue()).To(Equal("positional"))
		})

		It("treats words after -- as words even if they look like flags", func() {
			result, err := cli.Parse("-- --not-a-flag")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(2))
			Expect(result.GetTokens()[0].GetEndOfOptions()).NotTo(BeNil())
			Expect(result.GetTokens()[1].GetWord()).NotTo(BeNil())
			Expect(result.GetTokens()[1].GetWord().GetValue()).To(Equal("--not-a-flag"))
		})

		It("treats short-flag-looking token after -- as word", func() {
			result, err := cli.Parse("-- -v")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(2))
			Expect(result.GetTokens()[0].GetEndOfOptions()).NotTo(BeNil())
			Expect(result.GetTokens()[1].GetWord()).NotTo(BeNil())
			Expect(result.GetTokens()[1].GetWord().GetValue()).To(Equal("-v"))
		})

		It("treats short-flag cluster after -- as word", func() {
			result, err := cli.Parse("-- -abc")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(2))
			Expect(result.GetTokens()[0].GetEndOfOptions()).NotTo(BeNil())
			Expect(result.GetTokens()[1].GetWord()).NotTo(BeNil())
			Expect(result.GetTokens()[1].GetWord().GetValue()).To(Equal("-abc"))
		})

		It("treats -- after -- as a word", func() {
			result, err := cli.Parse("-- --")

			Expect(err).NotTo(HaveOccurred())
			Expect(result.GetTokens()).To(HaveLen(2))
			Expect(result.GetTokens()[0].GetEndOfOptions()).NotTo(BeNil())
			Expect(result.GetTokens()[1].GetWord()).NotTo(BeNil())
			Expect(result.GetTokens()[1].GetWord().GetValue()).To(Equal("--"))
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

	It("parses short flag followed by a word", func() {
		result, err := cli.Parse("-v input.txt")

		Expect(err).NotTo(HaveOccurred())
		Expect(result.GetTokens()).To(HaveLen(2))
		Expect(result.GetTokens()[0].GetShortFlag()).NotTo(BeNil())
		Expect(result.GetTokens()[0].GetShortFlag().GetNames()).To(Equal([]string{"v"}))
		Expect(result.GetTokens()[1].GetWord()).NotTo(BeNil())
		Expect(result.GetTokens()[1].GetWord().GetValue()).To(Equal("input.txt"))
	})

	It("parses short flag followed by a long flag", func() {
		result, err := cli.Parse("-v --output=file.txt")

		Expect(err).NotTo(HaveOccurred())
		Expect(result.GetTokens()).To(HaveLen(2))
		Expect(result.GetTokens()[0].GetShortFlag()).NotTo(BeNil())
		Expect(result.GetTokens()[0].GetShortFlag().GetNames()).To(Equal([]string{"v"}))
		Expect(result.GetTokens()[1].GetLongFlag()).NotTo(BeNil())
		Expect(result.GetTokens()[1].GetLongFlag().GetName()).To(Equal("output"))
	})

	It("parses mixed short flag, long flag with value, and word", func() {
		result, err := cli.Parse("-v --output=foo bar")

		Expect(err).NotTo(HaveOccurred())
		Expect(result.GetTokens()).To(HaveLen(3))
		Expect(result.GetTokens()[0].GetShortFlag()).NotTo(BeNil())
		Expect(result.GetTokens()[0].GetShortFlag().GetNames()).To(Equal([]string{"v"}))
		Expect(result.GetTokens()[1].GetLongFlag()).NotTo(BeNil())
		Expect(result.GetTokens()[1].GetLongFlag().GetName()).To(Equal("output"))
		Expect(result.GetTokens()[1].GetLongFlag().GetAssignment()).NotTo(BeNil())
		Expect(result.GetTokens()[1].GetLongFlag().GetAssignment().GetValue()).To(Equal("foo"))
		Expect(result.GetTokens()[2].GetWord()).NotTo(BeNil())
		Expect(result.GetTokens()[2].GetWord().GetValue()).To(Equal("bar"))
	})

	It("parses all four token types in sequence", func() {
		result, err := cli.Parse("cmd -v --output=foo -- bar")

		Expect(err).NotTo(HaveOccurred())
		Expect(result.GetTokens()).To(HaveLen(5))
		Expect(result.GetTokens()[0].GetWord()).NotTo(BeNil())
		Expect(result.GetTokens()[0].GetWord().GetValue()).To(Equal("cmd"))
		Expect(result.GetTokens()[1].GetShortFlag()).NotTo(BeNil())
		Expect(result.GetTokens()[1].GetShortFlag().GetNames()).To(Equal([]string{"v"}))
		Expect(result.GetTokens()[2].GetLongFlag()).NotTo(BeNil())
		Expect(result.GetTokens()[2].GetLongFlag().GetName()).To(Equal("output"))
		Expect(result.GetTokens()[2].GetLongFlag().GetAssignment().GetValue()).To(Equal("foo"))
		Expect(result.GetTokens()[3].GetEndOfOptions()).NotTo(BeNil())
		Expect(result.GetTokens()[4].GetWord()).NotTo(BeNil())
		Expect(result.GetTokens()[4].GetWord().GetValue()).To(Equal("bar"))
	})

	It("parses short flag cluster in mixed sequence", func() {
		result, err := cli.Parse("-abc --flag=val word")

		Expect(err).NotTo(HaveOccurred())
		Expect(result.GetTokens()).To(HaveLen(3))
		Expect(result.GetTokens()[0].GetShortFlag()).NotTo(BeNil())
		Expect(result.GetTokens()[0].GetShortFlag().GetNames()).To(Equal([]string{"a", "b", "c"}))
		Expect(result.GetTokens()[0].GetShortFlag().GetAssignment()).To(BeNil())
		Expect(result.GetTokens()[1].GetLongFlag()).NotTo(BeNil())
		Expect(result.GetTokens()[1].GetLongFlag().GetName()).To(Equal("flag"))
		Expect(result.GetTokens()[1].GetLongFlag().GetAssignment().GetValue()).To(Equal("val"))
		Expect(result.GetTokens()[2].GetWord()).NotTo(BeNil())
		Expect(result.GetTokens()[2].GetWord().GetValue()).To(Equal("word"))
	})
})
