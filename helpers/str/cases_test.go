package str

import (
	"github.com/davfer/archit/patterns/opts"
	"reflect"
	"testing"
)

func TestGetWords(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		ops    []opts.Opt[caseOpts]
		wantWs Words
		wantCs []Case
	}{
		{
			name:   "simple",
			in:     "simple",
			wantWs: Words{"simple"},
			wantCs: []Case{},
		},
		{
			name:   "empty",
			in:     "",
			wantWs: Words{""},
			wantCs: []Case{},
		},
		{
			name:   "snake",
			in:     "snake_case",
			wantWs: Words{"snake", "case"},
			wantCs: []Case{Snake},
		},
		{
			name:   "kebab",
			in:     "kebab-case",
			wantWs: Words{"kebab", "case"},
			wantCs: []Case{Kebab},
		},
		{
			name:   "pascal",
			in:     "PascalCase",
			wantWs: Words{"pascal", "case"},
			wantCs: []Case{Pascal},
		},
		{
			name:   "camel",
			in:     "camelCase",
			wantWs: Words{"camel", "case"},
			wantCs: []Case{Camel},
		},
		{
			name:   "whitespace",
			in:     "whitespace separated",
			wantWs: Words{"whitespace", "separated"},
			wantCs: []Case{Whitespace},
		},
		{
			name:   "user Separator",
			in:     "user|Separator|separated",
			ops:    []opts.Opt[caseOpts]{WithSeparator("|")},
			wantWs: Words{"user", "Separator", "separated"},
			wantCs: []Case{UserSeparator},
		},
		{
			name:   "ordering default",
			in:     "AWide_eyed RedFox-AAA-goesQuickly",
			wantWs: Words{"a", "wide", "eyed", "red", "fox", "a", "a", "a", "goes", "quickly"},
			wantCs: []Case{Whitespace, Kebab, Snake, Pascal, Camel},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWs, gotCs := GetWords(tt.in, tt.ops...)
			if !reflect.DeepEqual(gotWs, tt.wantWs) {
				t.Errorf("GetWords() gotWs = %+v, want %+v", gotWs, tt.wantWs)
			}
			if !reflect.DeepEqual(gotCs, tt.wantCs) {
				t.Errorf("GetWords() gotCs = %v, want %v", gotCs, tt.wantCs)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name string
		str  string
		from Case
		to   Case
		want string
	}{
		{
			name: "simple",
			str:  "simple",
			from: Camel,
			to:   Snake,
			want: "simple",
		},
		{
			name: "empty",
			str:  "",
			from: Camel,
			to:   Snake,
			want: "",
		},
		{
			name: "snake",
			str:  "snake_case",
			from: Snake,
			to:   Kebab,
			want: "snake-case",
		},
		{
			name: "kebab",
			str:  "kebab-case",
			from: Kebab,
			to:   Snake,
			want: "kebab_case",
		},
		{
			name: "pascal",
			str:  "PascalCase",
			from: Pascal,
			to:   Camel,
			want: "pascalCase",
		},
		{
			name: "camel",
			str:  "camelCase",
			from: Camel,
			to:   Pascal,
			want: "CamelCase",
		},
		{
			name: "whitespace",
			str:  "whitespace separated",
			from: Whitespace,
			to:   Snake,
			want: "whitespace_separated",
		},
		{
			name: "lower",
			str:  "Lower",
			from: Pascal,
			to:   Camel,
			want: "lower",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert(tt.str, tt.from, tt.to); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
