package main

import (
	"strings"
	"testing"
)

func TestNormalizeName(t *testing.T) {
	cases := map[string]string{
		"API Server":      "api-server",
		" Postgres ":      "postgres",
		"":                "unnamed",
		"User-Service_v2": "user-service-v2",
	}
	for in, want := range cases {
		got := NormalizeName(in)
		if got != want {
			t.Fatalf("NormalizeName(%q) = %q; want %q", in, got, want)
		}
	}
}

func TestGenerateMXGraphContainsNodes(t *testing.T) {
	apiID := NodeID("container", "API Server")
	dbID := NodeID("database", "Postgres")
	nodes := []Node{{ID: apiID, Label: "API Server"}, {ID: dbID, Label: "Postgres"}}
	edges := []Edge{{ID: EdgeID(apiID, dbID), Source: apiID, Target: dbID}}
	xml := GenerateMXGraph(nodes, edges)
	if !strings.Contains(xml, apiID) || !strings.Contains(xml, dbID) {
		t.Fatalf("generated xml missing node ids; got: %s", xml)
	}
	if !strings.Contains(xml, "mxGraphModel") {
		t.Fatalf("generated xml doesn't include mxGraphModel root")
	}
}
