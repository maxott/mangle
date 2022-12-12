// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package factstore

import (
	"fmt"
	"testing"

	"bitbucket.org/creachadair/stringset"
	"github.com/google/mangle/ast"
	"github.com/google/mangle/builtin"
	"github.com/google/mangle/parse"
)

func atom(s string) ast.Atom {
	term, err := parse.Term(s)
	if err != nil {
		panic(err)
	}
	return term.(ast.Atom)
}

func evalAtom(s string) ast.Atom {
	term, err := parse.Term(s)
	if err != nil {
		panic(err)
	}
	eval, err := builtin.EvalAtom(term.(ast.Atom), nil)
	if err != nil {
		panic(err)
	}
	return eval
}

func TestAddContains(t *testing.T) {
	for _, fs := range []FactStore{NewSimpleInMemoryStore(), NewIndexedInMemoryStore(), NewMultiIndexedInMemoryStore()} {
		t.Run(fmt.Sprintf("%T", fs), func(*testing.T) {
			tests := []ast.Atom{
				atom("baz()"),
				atom("foo(/bar)"),
				atom("foo(/zzz)"),
				atom("bar(/abc)"),
				atom("bar(/bar,/baz)"),
				atom("bar(/bar,/def)"),
				atom("bar(/abc,/def)"),
				evalAtom("bar([/abc],1,/def)"),
				evalAtom("baz([/abc : 1,  /def : 2], 1, /def)"),
				evalAtom("baz({/abc : 1,  /def : 2}, 1, /def)"),
			}
			for _, atom := range tests {
				if got := fs.Add(atom); !got {
					t.Errorf("for %v expected %v got %v", atom, true, got)
				}
				if !fs.Contains(atom) {
					t.Errorf("expected %v to be present store %v", atom, fs)
				}
				if got := fs.Add(atom); got {
					t.Errorf("for %v expected %v got %v", atom, false, got)
				}
			}

			for _, tt := range []struct {
				atom string
				want stringset.Set
			}{
				{
					atom: "baz()",
					want: stringset.New("baz()"),
				},
				{
					atom: "baz(X)",
					want: stringset.New(),
				},
				{
					atom: "baaaaz()",
					want: stringset.New(),
				},
				{
					atom: "foo(/bar)",
					want: stringset.New("foo(/bar)"),
				},
				{
					atom: "foo(/abc)",
					want: stringset.New(),
				},
				{
					atom: "fooooo(/bar)",
					want: stringset.New(),
				},
				{
					atom: "foo(X)",
					want: stringset.New("foo(/bar)", "foo(/zzz)"),
				},
				{
					atom: "bar(/bar,X)",
					want: stringset.New("bar(/bar,/baz)", "bar(/bar,/def)"),
				},
				{
					atom: "bar(X,Y)",
					want: stringset.New("bar(/bar,/baz)", "bar(/bar,/def)", "bar(/abc,/def)"),
				},
			} {
				t.Run(tt.atom, func(t *testing.T) {
					got := stringset.New()
					fs.GetFacts(atom(tt.atom), func(fact ast.Atom) error {
						got.Add(fact.String())
						return nil
					})
					if !got.Equals(tt.want) {
						t.Errorf("GetFacts(%q) = %v want %v", tt.atom, got, tt.want)
					}
				})
			}

			if got, want := fs.EstimateFactCount(), len(tests); got != want {
				t.Errorf("EstimateFactCount() = %d want %d", got, want)
			}
		})
	}
}

func TestTeeingAddContainsMerge(t *testing.T) {
	base := NewSimpleInMemoryStore()
	base.Add(atom("foo(/bar)"))

	tmp := NewTeeingStore(base)
	tmp.Add(atom("foo(/zzz)"))

	if base.Contains(atom("foo(/zzz)")) {
		t.Errorf("base.Contains(%v) = true want false", atom("foo(/zzz)"))
	}
	if !tmp.Contains(atom("foo(/zzz)")) {
		t.Errorf("tmp.Contains(%v) = false want true", atom("foo(/zzz)"))
	}
	if !tmp.Contains(atom("foo(/bar)")) {
		t.Errorf("tmp.Contains(%v) = false want true", atom("foo(/bar)"))
	}
	if tmp.Out.Contains(atom("foo(/bar)")) {
		t.Errorf("tmp.Out.Contains(%v) = true want false", atom("foo(/bar)"))
	}
	if got, want := base.EstimateFactCount(), 1; got != want {
		t.Errorf("base.EstimateFactCount() = %d want %d", got, want)
	}
	if got, want := tmp.EstimateFactCount(), 2; got != want {
		t.Errorf("tmp.EstimateFactCount() = %d want %d", got, want)
	}

	tmp.Merge(base)
	if !tmp.Out.Contains(atom("foo(/bar)")) {
		t.Errorf("tmp.Out.Contains(%v) = false want true", atom("foo(/bar)"))
	}

	if got, want := tmp.EstimateFactCount(), 2; got < want {
		t.Errorf("tmp.EstimateFactCount() = %d want at least %d", got, want)
	}

}
