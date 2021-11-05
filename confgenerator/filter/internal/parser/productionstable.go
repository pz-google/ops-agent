// Code generated by gocc; DO NOT EDIT.

package parser

// Copyright 2021 Google LLC
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

import (
	"github.com/GoogleCloudPlatform/ops-agent/confgenerator/filter/internal/ast"
	"github.com/GoogleCloudPlatform/ops-agent/confgenerator/filter/internal/token"
)

type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib, interface{}) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Filter	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Filter : Expression	<< ast.Simplify(X[0]) >>`,
		Id:         "Filter",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.Simplify(X[0])
		},
	},
	ProdTabEntry{
		String: `Filter : ws Expression	<< ast.Simplify(X[1]) >>`,
		Id:         "Filter",
		NTType:     1,
		Index:      2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.Simplify(X[1])
		},
	},
	ProdTabEntry{
		String: `Filter : Expression ws	<< ast.Simplify(X[0]) >>`,
		Id:         "Filter",
		NTType:     1,
		Index:      3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.Simplify(X[0])
		},
	},
	ProdTabEntry{
		String: `Filter : ws Expression ws	<< ast.Simplify(X[1]) >>`,
		Id:         "Filter",
		NTType:     1,
		Index:      4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.Simplify(X[1])
		},
	},
	ProdTabEntry{
		String: `Expression : AmbiguousSequence	<< ast.NewConjunction(X[0]) >>`,
		Id:         "Expression",
		NTType:     2,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewConjunction(X[0])
		},
	},
	ProdTabEntry{
		String: `Expression : Expression andOp AmbiguousSequence	<< X[0].(ast.Conjunction).Append(X[2]) >>`,
		Id:         "Expression",
		NTType:     2,
		Index:      6,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0].(ast.Conjunction).Append(X[2])
		},
	},
	ProdTabEntry{
		String: `AmbiguousSequence : AmbiguousFactor	<< ast.NewConjunction(X[0]) >>`,
		Id:         "AmbiguousSequence",
		NTType:     3,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewConjunction(X[0])
		},
	},
	ProdTabEntry{
		String: `AmbiguousSequence : AmbiguousSequence ws AmbiguousFactor	<< X[0].(ast.Conjunction).Append(X[2]) >>`,
		Id:         "AmbiguousSequence",
		NTType:     3,
		Index:      8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0].(ast.Conjunction).Append(X[2])
		},
	},
	ProdTabEntry{
		String: `AmbiguousFactor : Term	<< ast.NewDisjunction(X[0]) >>`,
		Id:         "AmbiguousFactor",
		NTType:     4,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewDisjunction(X[0])
		},
	},
	ProdTabEntry{
		String: `AmbiguousFactor : AmbiguousFactor orOp Term	<< X[0].(ast.Disjunction).Append(X[2]) >>`,
		Id:         "AmbiguousFactor",
		NTType:     4,
		Index:      10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0].(ast.Disjunction).Append(X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Primitive	<<  >>`,
		Id:         "Term",
		NTType:     5,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : not Primitive	<< &ast.Negation{X[1].(ast.Expression)}, nil >>`,
		Id:         "Term",
		NTType:     5,
		Index:      12,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return &ast.Negation{X[1].(ast.Expression)}, nil
		},
	},
	ProdTabEntry{
		String: `Term : not ws Primitive	<< &ast.Negation{X[2].(ast.Expression)}, nil >>`,
		Id:         "Term",
		NTType:     5,
		Index:      13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return &ast.Negation{X[2].(ast.Expression)}, nil
		},
	},
	ProdTabEntry{
		String: `Term : minus Primitive	<< &ast.Negation{X[1].(ast.Expression)}, nil >>`,
		Id:         "Term",
		NTType:     5,
		Index:      14,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return &ast.Negation{X[1].(ast.Expression)}, nil
		},
	},
	ProdTabEntry{
		String: `Primitive : Restriction	<<  >>`,
		Id:         "Primitive",
		NTType:     6,
		Index:      15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Primitive : Composite	<<  >>`,
		Id:         "Primitive",
		NTType:     6,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Restriction : Comparable	<< ast.NewRestriction(X[0], "GLOBAL", nil) >>`,
		Id:         "Restriction",
		NTType:     7,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewRestriction(X[0], "GLOBAL", nil)
		},
	},
	ProdTabEntry{
		String: `Restriction : Comparable Comparator Arg	<< ast.NewRestriction(X[0], X[1], X[2]) >>`,
		Id:         "Restriction",
		NTType:     7,
		Index:      18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewRestriction(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `Comparable : Member	<<  >>`,
		Id:         "Comparable",
		NTType:     8,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Member : Item	<< ast.Member{X[0].(string)}, nil >>`,
		Id:         "Member",
		NTType:     9,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.Member{X[0].(string)}, nil
		},
	},
	ProdTabEntry{
		String: `Member : Member dot ItemKeyword	<< append(X[0].(ast.Member), X[2].(string)), nil >>`,
		Id:         "Member",
		NTType:     9,
		Index:      21,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return append(X[0].(ast.Member), X[2].(string)), nil
		},
	},
	ProdTabEntry{
		String: `Composite : lparen Expression rparen	<< X[1], nil >>`,
		Id:         "Composite",
		NTType:     10,
		Index:      22,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Composite : lparen Expression ws rparen	<< X[1], nil >>`,
		Id:         "Composite",
		NTType:     10,
		Index:      23,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Composite : lparen ws Expression rparen	<< X[2], nil >>`,
		Id:         "Composite",
		NTType:     10,
		Index:      24,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `Composite : lparen ws Expression ws rparen	<< X[2], nil >>`,
		Id:         "Composite",
		NTType:     10,
		Index:      25,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `Arg : Comparable	<<  >>`,
		Id:         "Arg",
		NTType:     11,
		Index:      26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Item : Value	<<  >>`,
		Id:         "Item",
		NTType:     12,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Value : text	<< ast.ParseText(X[0]) >>`,
		Id:         "Value",
		NTType:     13,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.ParseText(X[0])
		},
	},
	ProdTabEntry{
		String: `Value : Phrase	<<  >>`,
		Id:         "Value",
		NTType:     13,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Phrase : string	<< ast.ParseString(X[0]) >>`,
		Id:         "Phrase",
		NTType:     14,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.ParseString(X[0])
		},
	},
	ProdTabEntry{
		String: `ItemKeyword : Item	<<  >>`,
		Id:         "ItemKeyword",
		NTType:     15,
		Index:      31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ItemKeyword : Keyword	<< string(X[0].(*token.Token).Lit), nil >>`,
		Id:         "ItemKeyword",
		NTType:     15,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return string(X[0].(*token.Token).Lit), nil
		},
	},
	ProdTabEntry{
		String: `Keyword : or	<<  >>`,
		Id:         "Keyword",
		NTType:     16,
		Index:      33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Keyword : and	<<  >>`,
		Id:         "Keyword",
		NTType:     16,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Keyword : not	<<  >>`,
		Id:         "Keyword",
		NTType:     16,
		Index:      35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comparator : Comparison	<<  >>`,
		Id:         "Comparator",
		NTType:     17,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comparator : ws Comparison	<< X[1], nil >>`,
		Id:         "Comparator",
		NTType:     17,
		Index:      37,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Comparator : Comparison ws	<<  >>`,
		Id:         "Comparator",
		NTType:     17,
		Index:      38,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comparator : ws Comparison ws	<< X[1], nil >>`,
		Id:         "Comparator",
		NTType:     17,
		Index:      39,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Comparison : less_equals	<<  >>`,
		Id:         "Comparison",
		NTType:     18,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comparison : less_than	<<  >>`,
		Id:         "Comparison",
		NTType:     18,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comparison : greater_equals	<<  >>`,
		Id:         "Comparison",
		NTType:     18,
		Index:      42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comparison : greater_than	<<  >>`,
		Id:         "Comparison",
		NTType:     18,
		Index:      43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comparison : not_equals	<<  >>`,
		Id:         "Comparison",
		NTType:     18,
		Index:      44,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comparison : equals	<<  >>`,
		Id:         "Comparison",
		NTType:     18,
		Index:      45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comparison : has	<<  >>`,
		Id:         "Comparison",
		NTType:     18,
		Index:      46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comparison : matches_regexp	<<  >>`,
		Id:         "Comparison",
		NTType:     18,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comparison : not_matches_regexp	<<  >>`,
		Id:         "Comparison",
		NTType:     18,
		Index:      48,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
}
