package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

// Node represents a diagram node
type Node struct {
	ID    string
	Label string
	Style string
	X     int
	Y     int
	W     int
	H     int
}

// Edge represents a diagram edge
type Edge struct {
	ID     string
	Source string
	Target string
	Label  string
	Style  string
}

// NormalizeName converts a label to a deterministic id fragment
func NormalizeName(s string) string {
	s = strings.TrimSpace(s)
	// NFC and unicode handling omitted for brevity
	s = strings.ToLower(s)
	// replace runs of non-alnum with -
	re := regexp.MustCompile(`[^a-z0-9]+`)
	s = re.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	if s == "" {
		return "unnamed"
	}
	return s
}

// NodeID returns deterministic node id for a type and name
func NodeID(typ, name string) string {
	return fmt.Sprintf("node://%s/%s", typ, NormalizeName(name))
}

// EdgeID returns deterministic edge id between two nodes
func EdgeID(fromID, toID string) string {
	// keep edge ids reasonably short
	f := strings.ReplaceAll(fromID, "://", "--")
	t := strings.ReplaceAll(toID, "://", "--")
	return fmt.Sprintf("edge://%s-to-%s", f, t)
}

// GenerateMXGraph produces a minimal mxfile XML string for nodes and edges
func GenerateMXGraph(nodes []Node, edges []Edge) string {
	var b bytes.Buffer
	b.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	b.WriteString("<mxfile host=\"app.diagrams.net\">\n")
	b.WriteString("  <diagram id=\"diagram-1\" name=\"Page-1\">\n")
	b.WriteString("    <mxGraphModel dx=\"1000\" dy=\"800\" grid=\"1\" gridSize=\"10\" guides=\"1\" tooltips=\"1\" connect=\"1\" arrows=\"1\">\n")
	b.WriteString("      <root>\n")
	b.WriteString("        <mxCell id=\"0\"/>\n")
	b.WriteString("        <mxCell id=\"1\" parent=\"0\"/>\n")

	for _, n := range nodes {
		style := n.Style
		if style == "" {
			style = "rounded=1;fillColor=#0b5fff;fontColor=#ffffff;"
		}
		b.WriteString(fmt.Sprintf("        <mxCell id=\"%s\" value=\"%s\" style=\"%s\" vertex=\"1\" parent=\"1\">\n", n.ID, escape(n.Label), style))
		b.WriteString(fmt.Sprintf("          <mxGeometry x=\"%d\" y=\"%d\" width=\"%d\" height=\"%d\" as=\"geometry\"/>\n", n.X, n.Y, n.W, n.H))
		b.WriteString("        </mxCell>\n")
	}

	for _, e := range edges {
		style := e.Style
		if style == "" {
			style = "edgeStyle=elbowEdgeStyle;rounded=0;"
		}
		b.WriteString(fmt.Sprintf("        <mxCell id=\"%s\" style=\"%s\" edge=\"1\" parent=\"1\" source=\"%s\" target=\"%s\">\n", e.ID, style, e.Source, e.Target))
		b.WriteString("          <mxGeometry relative=\"1\" as=\"geometry\"/>\n")
		b.WriteString("        </mxCell>\n")
	}

	b.WriteString("      </root>\n")
	b.WriteString("    </mxGraphModel>\n")
	b.WriteString("  </diagram>\n")
	b.WriteString("</mxfile>\n")
	return b.String()
}

// escape minimal XML special chars for values
func escape(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	return s
}
