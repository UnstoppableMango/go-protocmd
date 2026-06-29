package cli

import (
	"strings"

	cliv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/unmango/cli/v1alpha1"
)

func Parse(cmdline string) (*cliv1alpha1.CommandLine, error) {
	b := &cliv1alpha1.CommandLine_builder{
		Raw: &cmdline,
	}

	for x := range strings.FieldsSeq(cmdline) {
		if tok, err := parseToken(x); err != nil {
			return nil, err
		} else {
			b.Tokens = append(b.Tokens, tok)
		}
	}

	return b.Build(), nil
}

func parseToken(tok string) (*cliv1alpha1.Token, error) {
	var err error
	b := &cliv1alpha1.Token_builder{}
	if strings.HasPrefix(tok, "--") {
		b.LongFlag, err = parseLongFlag(tok)
	} else if strings.HasPrefix(tok, "-") {
		b.ShortFlag, err = parseShortFlag(tok)
	} else {
		b.Word, err = parseWord(tok)
	}
	if err != nil {
		return nil, err
	}
	return b.Build(), nil
}

func parseLongFlag(tok string) (*cliv1alpha1.LongFlag, error) {
	b := &cliv1alpha1.LongFlag_builder{
		Name: new(strings.TrimPrefix(tok, "--")),
	}
	return b.Build(), nil
}

func parseShortFlag(tok string) (*cliv1alpha1.ShortFlag, error) {
	panic("not implemented")
}

func parseWord(tok string) (*cliv1alpha1.Word, error) {
	panic("not implemented")
}
