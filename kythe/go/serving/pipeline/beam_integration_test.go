/*
 * Copyright 2018 The Kythe Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pipeline

import (
	"context"
	"strings"
	"testing"

	"kythe.io/kythe/go/services/xrefs"
	xsrv "kythe.io/kythe/go/serving/xrefs"
	"kythe.io/kythe/go/storage/inmemory"
	"kythe.io/kythe/go/storage/keyvalue"
	"kythe.io/kythe/go/util/kytheuri"

	"github.com/apache/beam/sdks/go/pkg/beam"
	"github.com/apache/beam/sdks/go/pkg/beam/testing/ptest"
	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"

	cpb "kythe.io/kythe/proto/common_go_proto"
	scpb "kythe.io/kythe/proto/schema_go_proto"
	spb "kythe.io/kythe/proto/storage_go_proto"
	xpb "kythe.io/kythe/proto/xref_go_proto"
)

var ctx = context.Background()

func encodeMarkedSource(ms *cpb.MarkedSource) []byte {
	rec, err := proto.Marshal(ms)
	if err != nil {
		panic(err)
	}
	return rec
}

func TestServingSimpleDecorations(t *testing.T) {
	file := &spb.VName{Path: "path"}
	const expectedText = "some text\n"
	testNodes := []*scpb.Node{{
		Source: file,
		Kind:   &scpb.Node_KytheKind{scpb.NodeKind_FILE},
		Fact: []*scpb.Fact{{
			Name:  &scpb.Fact_KytheName{scpb.FactName_TEXT},
			Value: []byte(expectedText),
		}, {
			Name:  &scpb.Fact_KytheName{scpb.FactName_TEXT_ENCODING},
			Value: []byte("ascii"),
		}},
	}, {
		Source: &spb.VName{Path: "path", Signature: "anchor1"},
		Kind:   &scpb.Node_KytheKind{scpb.NodeKind_ANCHOR},
		Fact: []*scpb.Fact{{
			Name:  &scpb.Fact_KytheName{scpb.FactName_LOC_START},
			Value: []byte("0"),
		}, {
			Name:  &scpb.Fact_KytheName{scpb.FactName_LOC_END},
			Value: []byte("4"),
		}},
		Edge: []*scpb.Edge{{
			Kind:   &scpb.Edge_KytheKind{scpb.EdgeKind_REF},
			Target: &spb.VName{Signature: "simpleDecor"},
		}},
	}, {
		Source: &spb.VName{Signature: "simpleDecor"},
		Kind:   &scpb.Node_KytheKind{scpb.NodeKind_RECORD},
	}, {
		Source: &spb.VName{Path: "path", Signature: "anchor2"},
		Kind:   &scpb.Node_KytheKind{scpb.NodeKind_ANCHOR},
		Fact: []*scpb.Fact{{
			Name:  &scpb.Fact_KytheName{scpb.FactName_LOC_START},
			Value: []byte("5"),
		}, {
			Name:  &scpb.Fact_KytheName{scpb.FactName_LOC_END},
			Value: []byte("9"),
		}},
		Edge: []*scpb.Edge{{
			Kind:   &scpb.Edge_KytheKind{scpb.EdgeKind_REF},
			Target: &spb.VName{Signature: "decorWithDef"},
		}},
	}, {
		Source: &spb.VName{Signature: "def1"},
		Kind:   &scpb.Node_KytheKind{scpb.NodeKind_ANCHOR},
		Fact: []*scpb.Fact{{
			Name:  &scpb.Fact_KytheName{scpb.FactName_LOC_START},
			Value: []byte("0"),
		}, {
			Name:  &scpb.Fact_KytheName{scpb.FactName_LOC_END},
			Value: []byte("3"),
		}},
		Edge: []*scpb.Edge{{
			Kind:   &scpb.Edge_KytheKind{scpb.EdgeKind_DEFINES},
			Target: &spb.VName{Signature: "decorWithDef"},
		}},
	}, {
		Source: &spb.VName{},
		Kind:   &scpb.Node_KytheKind{scpb.NodeKind_FILE},
		Fact: []*scpb.Fact{{
			Name:  &scpb.Fact_KytheName{scpb.FactName_TEXT},
			Value: []byte("def\n"),
		}},
	}}

	p, s, nodes := ptest.CreateList(testNodes)
	decor := FromNodes(s, nodes).SplitDecorations()

	db := inmemory.NewKeyValueDB()
	w, err := db.Writer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Mark table as columnar
	if err := w.Write([]byte(xsrv.ColumnarTableKeyMarker), []byte{}); err != nil {
		t.Fatal(err)
	}
	// Write columnar data to inmemory.KeyValueDB
	beam.ParDo(s, &writeTo{w}, decor)

	if err := ptest.Run(p); err != nil {
		t.Fatalf("Pipeline error: %+v", err)
	} else if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	xs := xsrv.NewService(ctx, db)
	fileTicket := kytheuri.ToString(file)

	t.Run("source_text", makeDecorTestCase(ctx, xs, &xpb.DecorationsRequest{
		Location:   &xpb.Location{Ticket: fileTicket},
		SourceText: true,
	}, &xpb.DecorationsReply{
		Location:   &xpb.Location{Ticket: fileTicket},
		SourceText: []byte(expectedText),
		Encoding:   "ascii",
	}))

	t.Run("references", makeDecorTestCase(ctx, xs, &xpb.DecorationsRequest{
		Location:   &xpb.Location{Ticket: fileTicket},
		References: true,
	}, &xpb.DecorationsReply{
		Location: &xpb.Location{Ticket: fileTicket},
		Reference: []*xpb.DecorationsReply_Reference{{
			Span: &cpb.Span{
				Start: &cpb.Point{
					LineNumber: 1,
				},
				End: &cpb.Point{
					ByteOffset:   4,
					ColumnOffset: 4,
					LineNumber:   1,
				},
			},
			Kind:         "/kythe/edge/ref",
			TargetTicket: "kythe:#simpleDecor",
		}, {
			Span: &cpb.Span{
				Start: &cpb.Point{
					ByteOffset:   5,
					ColumnOffset: 5,
					LineNumber:   1,
				},
				End: &cpb.Point{
					ByteOffset:   9,
					ColumnOffset: 9,
					LineNumber:   1,
				},
			},
			Kind:         "/kythe/edge/ref",
			TargetTicket: "kythe:#decorWithDef",
			// TargetDefinition: explicitly not requested
		}},
		// Nodes: not requested
		// DefinitionLocations: not requested
	}))

	t.Run("referenced_nodes", makeDecorTestCase(ctx, xs, &xpb.DecorationsRequest{
		Location:   &xpb.Location{Ticket: fileTicket},
		References: true,
		Filter:     []string{"**"},
	}, &xpb.DecorationsReply{
		Location: &xpb.Location{Ticket: fileTicket},
		Reference: []*xpb.DecorationsReply_Reference{{
			Span: &cpb.Span{
				Start: &cpb.Point{
					LineNumber: 1,
				},
				End: &cpb.Point{
					ByteOffset:   4,
					ColumnOffset: 4,
					LineNumber:   1,
				},
			},
			Kind:         "/kythe/edge/ref",
			TargetTicket: "kythe:#simpleDecor",
		}, {
			Span: &cpb.Span{
				Start: &cpb.Point{
					ByteOffset:   5,
					ColumnOffset: 5,
					LineNumber:   1,
				},
				End: &cpb.Point{
					ByteOffset:   9,
					ColumnOffset: 9,
					LineNumber:   1,
				},
			},
			Kind:         "/kythe/edge/ref",
			TargetTicket: "kythe:#decorWithDef",
			// TargetDefinition: explicitly not requested
		}},
		Nodes: map[string]*cpb.NodeInfo{
			"kythe:#simpleDecor": {
				Facts: map[string][]byte{
					"/kythe/node/kind": []byte("record"),
				},
			},
		},
		// DefinitionLocations: not requested
	}))

	t.Run("target_definitions", makeDecorTestCase(ctx, xs, &xpb.DecorationsRequest{
		Location:          &xpb.Location{Ticket: fileTicket},
		References:        true,
		TargetDefinitions: true,
	}, &xpb.DecorationsReply{
		Location: &xpb.Location{Ticket: fileTicket},
		Reference: []*xpb.DecorationsReply_Reference{{
			Span: &cpb.Span{
				Start: &cpb.Point{
					LineNumber: 1,
				},
				End: &cpb.Point{
					ByteOffset:   4,
					ColumnOffset: 4,
					LineNumber:   1,
				},
			},
			Kind:         "/kythe/edge/ref",
			TargetTicket: "kythe:#simpleDecor",
		}, {
			Span: &cpb.Span{
				Start: &cpb.Point{
					ByteOffset:   5,
					ColumnOffset: 5,
					LineNumber:   1,
				},
				End: &cpb.Point{
					ByteOffset:   9,
					ColumnOffset: 9,
					LineNumber:   1,
				},
			},
			Kind:             "/kythe/edge/ref",
			TargetTicket:     "kythe:#decorWithDef",
			TargetDefinition: "kythe:#def1", // expected definition
		}},
		// Nodes: not requested
		DefinitionLocations: map[string]*xpb.Anchor{
			"kythe:#def1": {
				Ticket: "kythe:#def1",
				Parent: "kythe:",
				Span: &cpb.Span{
					Start: &cpb.Point{
						LineNumber: 1,
					},
					End: &cpb.Point{
						ByteOffset:   3,
						ColumnOffset: 3,
						LineNumber:   1,
					},
				},
				Snippet: "def",
				SnippetSpan: &cpb.Span{
					Start: &cpb.Point{
						LineNumber: 1,
					},
					End: &cpb.Point{
						ByteOffset:   3,
						ColumnOffset: 3,
						LineNumber:   1,
					},
				},
			},
		},
	}))

	// TODO(schroederc): test split file contents
	// TODO(schroederc): test overrides
	// TODO(schroederc): test diagnostics (w/ or w/o span)
}

func TestServingSimpleCrossReferences(t *testing.T) {
	src := &spb.VName{Path: "path", Signature: "signature"}
	ms := &cpb.MarkedSource{
		Kind:    cpb.MarkedSource_IDENTIFIER,
		PreText: "identifier",
	}
	testNodes := []*scpb.Node{{
		Source: &spb.VName{Path: "path"},
		Kind:   &scpb.Node_KytheKind{scpb.NodeKind_FILE},
		Fact: []*scpb.Fact{{
			Name:  &scpb.Fact_KytheName{scpb.FactName_TEXT},
			Value: []byte("blah blah\n"),
		}, {
			Name:  &scpb.Fact_KytheName{scpb.FactName_TEXT_ENCODING},
			Value: []byte("ascii"),
		}},
	}, {
		Source: &spb.VName{Path: "path", Signature: "anchor1"},
		Kind:   &scpb.Node_KytheKind{scpb.NodeKind_ANCHOR},
		Fact: []*scpb.Fact{{
			Name:  &scpb.Fact_KytheName{scpb.FactName_LOC_START},
			Value: []byte("5"),
		}, {
			Name:  &scpb.Fact_KytheName{scpb.FactName_LOC_END},
			Value: []byte("9"),
		}},
		Edge: []*scpb.Edge{{
			Kind:   &scpb.Edge_KytheKind{scpb.EdgeKind_REF},
			Target: src,
		}},
	}, {
		Source: src,
		Kind:   &scpb.Node_KytheKind{scpb.NodeKind_RECORD},
		Fact: []*scpb.Fact{{
			Name:  &scpb.Fact_KytheName{scpb.FactName_CODE},
			Value: encodeMarkedSource(ms),
		}},
	}}

	p, s, nodes := ptest.CreateList(testNodes)
	xrefs := FromNodes(s, nodes).SplitCrossReferences()

	db := inmemory.NewKeyValueDB()
	w, err := db.Writer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Mark table as columnar
	if err := w.Write([]byte(xsrv.ColumnarTableKeyMarker), []byte{}); err != nil {
		t.Fatal(err)
	}
	// Write columnar data to inmemory.KeyValueDB
	beam.ParDo(s, &writeTo{w}, xrefs)

	if err := ptest.Run(p); err != nil {
		t.Fatalf("Pipeline error: %+v", err)
	} else if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	xs := xsrv.NewService(ctx, db)

	ticket := kytheuri.ToString(src)

	t.Run("requested_node", makeXRefTestCase(ctx, xs, &xpb.CrossReferencesRequest{
		Ticket:          []string{ticket},
		Filter:          []string{"**"},
		RelatedNodeKind: []string{"NONE"},
	}, &xpb.CrossReferencesReply{
		CrossReferences: map[string]*xpb.CrossReferencesReply_CrossReferenceSet{
			ticket: {
				Ticket:       ticket,
				MarkedSource: ms,
			},
		},
		Nodes: map[string]*cpb.NodeInfo{
			ticket: {
				Facts: map[string][]byte{
					"/kythe/node/kind": []byte("record"),

					// TODO(schroederc): ellide; MarkedSource already included
					"/kythe/code": encodeMarkedSource(ms),
				},
			},
		},
	}))

	t.Run("refs", makeXRefTestCase(ctx, xs, &xpb.CrossReferencesRequest{
		Ticket:        []string{ticket},
		ReferenceKind: xpb.CrossReferencesRequest_ALL_REFERENCES,
	}, &xpb.CrossReferencesReply{
		CrossReferences: map[string]*xpb.CrossReferencesReply_CrossReferenceSet{
			ticket: {
				Ticket:       ticket,
				MarkedSource: ms,
				Reference: []*xpb.CrossReferencesReply_RelatedAnchor{{
					Anchor: &xpb.Anchor{
						Parent: "kythe:?path=path",
						Span: &cpb.Span{
							Start: &cpb.Point{
								ByteOffset:   5,
								ColumnOffset: 5,
								LineNumber:   1,
							},
							End: &cpb.Point{
								ByteOffset:   9,
								ColumnOffset: 9,
								LineNumber:   1,
							},
						},
						Snippet: "blah blah",
						SnippetSpan: &cpb.Span{
							Start: &cpb.Point{
								LineNumber: 1,
							},
							End: &cpb.Point{
								ByteOffset:   9,
								ColumnOffset: 9,
								LineNumber:   1,
							},
						},
					},
				}},
			},
		},
	}))
}

type writeTo struct{ w keyvalue.Writer }

func (p *writeTo) ProcessElement(ctx context.Context, k, v []byte, emit func([]byte)) error {
	return p.w.Write(k, v)
}

func makeDecorTestCase(ctx context.Context, xs xrefs.Service, req *xpb.DecorationsRequest, expected *xpb.DecorationsReply) func(*testing.T) {
	return func(t *testing.T) {
		reply, err := xs.Decorations(ctx, req)
		if err != nil {
			t.Fatalf("Decorations error: %v", err)
		}
		if diff := cmp.Diff(expected, reply, ignoreProtoXXXFields); diff != "" {
			t.Fatalf("DecorationsReply differences: (- expected; + found)\n%s", diff)
		}
	}
}

func makeXRefTestCase(ctx context.Context, xs xrefs.Service, req *xpb.CrossReferencesRequest, expected *xpb.CrossReferencesReply) func(*testing.T) {
	return func(t *testing.T) {
		reply, err := xs.CrossReferences(ctx, req)
		if err != nil {
			t.Fatalf("CrossReferences error: %v", err)
		}
		if diff := cmp.Diff(expected, reply, ignoreProtoXXXFields); diff != "" {
			t.Fatalf("CrossReferencesReply differences: (- expected; + found)\n%s", diff)
		}
	}
}

var ignoreProtoXXXFields = cmp.FilterPath(func(p cmp.Path) bool {
	for _, s := range p {
		if strings.HasPrefix(s.String(), ".XXX_") {
			return true
		}
	}
	return false
}, cmp.Ignore())
