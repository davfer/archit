package str

import (
	"github.com/davfer/archit/helpers/sli"
	"github.com/davfer/archit/patterns/opts"
	"github.com/davfer/archit/patterns/parser"
	"strings"
)

var defaultEntries = []caseEntry{
	{c: Whitespace, parser: Separator{s: " "}},
	{c: Kebab, parser: Separator{s: "-"}},
	{c: Snake, parser: Separator{s: "_"}},
	{c: Pascal, parser: capitalCase{mustStartsWithUpper: true}},
	{c: Camel, parser: capitalCase{mustStartsWithUpper: false}},
}

type Case string

const (
	UserSeparator Case = "Separator"  // A Custom UserSeparator Separated String
	Whitespace    Case = "whitespace" // A Whitespace Separated String
	Kebab         Case = "kebab"      // A Kebab kebab-case-separated-string
	Snake         Case = "snake"      // A Snake snake_case_separated_string
	Camel         Case = "camel"      // A Camel camelCaseSeparatedString
	Pascal        Case = "pascal"     // A Pascal PascalCaseSeparatedString
)

type Words []string

type caseOpts struct {
	separators []string
	parsers    []caseEntry
}

type caseEntry struct {
	c      Case
	parser parser.Parser[Words, string]
}

func WithSeparator(separator string) opts.Opt[caseOpts] {
	return func(o caseOpts) caseOpts {
		o.separators = append(o.separators, separator)
		return o
	}
}

// GetWords returns the words and the cases of a string
func GetWords(s string, o ...opts.Opt[caseOpts]) (ws Words, cs []Case) {
	cfg := opts.New[caseOpts](o...)

	if cfg.parsers == nil {
		cfg.parsers = defaultEntries
	}

	if cfg.separators != nil {
		for _, sep := range cfg.separators {
			cfg.parsers = append([]caseEntry{{c: UserSeparator, parser: Separator{s: sep}}}, cfg.parsers...)
		}
	}

	ws, cs = parseWords(Words{s}, cfg.parsers)

	return
}

func parseWords(ws Words, ps []caseEntry) (Words, []Case) {
	if len(ps) == 0 {
		return ws, []Case{}
	}

	var p caseEntry
	p, ps = sli.PopHeap(ps)

	var resWords Words
	resCases := make([]Case, 0)
	for _, word := range ws {
		wsAux, okAux := p.parser.Parse(word)
		if okAux {
			resCases = sli.AppendIfMissing(resCases, p.c)
			resWords = append(resWords, wsAux...)
		} else {
			resWords = append(resWords, word)
		}
	}

	immWords, immCases := parseWords(resWords, ps)
	resCases = sli.AppendIfMissing(resCases, immCases...)

	return immWords, resCases
}

// Convert a string from one case to another
func Convert(str string, from, to Case) string {
	pf, okf := sli.Find(defaultEntries, func(p caseEntry) bool {
		return p.c == from
	})
	pt, okt := sli.Find(defaultEntries, func(p caseEntry) bool {
		return p.c == to
	})

	if !okf || !okt {
		return str
	}

	ws, ok := pf.parser.Parse(str)
	if !ok {
		return str
	}

	resStr, resOk := pt.parser.Build(ws)
	if !resOk {
		return str
	}

	return resStr
}

type capitalCase struct {
	mustStartsWithUpper bool
}

func (p capitalCase) Parse(s string) (ws Words, ok bool) {
	if len(s) == 0 {
		return
	}
	first := s[0]
	if p.mustStartsWithUpper && (first < 'A' || first > 'Z') {
		return
	}
	if !p.mustStartsWithUpper && (first < 'a' || first > 'z') {
		return
	}

	rest := s[1:]
	if strings.ToLower(rest) == rest {
		return
	}

	var column strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			column.WriteRune('_')
		}
		column.WriteRune(r)
	}

	ws = strings.Split(strings.ToLower(column.String()), "_")

	for i, w := range ws {
		if len(w) == 0 {
			ws = append(ws[:i], ws[i+1:]...)
		}
	}

	if len(ws) > 0 {
		ok = true
	}

	return
}

func (p capitalCase) Build(ws Words) (string, bool) {
	if len(ws) == 0 {
		return "", false
	}
	var column strings.Builder
	for _, w := range ws {
		if w == "" {
			continue
		}

		if p.mustStartsWithUpper {
			column.WriteString(strings.Title(w))
		} else {
			column.WriteString(w)
		}
	}

	return column.String(), true
}

type Separator struct {
	s string
}

func (s Separator) Parse(str string) (ws Words, ok bool) {
	if strings.Contains(str, s.s) {
		ws = strings.Split(str, s.s)
		ok = true
	}
	return
}

func (s Separator) Build(ws Words) (string, bool) {
	if len(ws) == 0 {
		return "", false
	}
	return strings.Join(ws, s.s), true
}
