package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	out := flag.String("out", "", "output .drawio file path (auto-generated if empty)")
	c4doc := flag.String("c4doc", "", "path to C4 markdown document")
	level := flag.String("level", "", "C4 level: L1, L2, L3, or L4")
	theme := flag.String("theme", "light", "diagram theme: light, dark, neutral")
	layout := flag.String("layout", "vertical", "layout: horizontal, vertical, layered")
	dry := flag.Bool("dry-run", false, "dry run (print summary instead of writing)")
	flag.Parse()

	// Validate inputs
	if *c4doc == "" && *level == "" {
		fmt.Fprintln(os.Stderr, "Error: must specify either --c4doc or --level")
		flag.Usage()
		os.Exit(2)
	}

	var doc *C4Doc
	var sources []string
	var err error

	// Parse C4 document if provided
	if *c4doc != "" {
		doc, err = ParseC4Doc(*c4doc)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing C4 doc: %v\n", err)
			os.Exit(1)
		}
		sources = append(sources, *c4doc)
		if *level == "" {
			*level = doc.Level
		}
	} else {
		// Create minimal doc from level
		doc = &C4Doc{Level: *level}
	}

	// Generate output filename if not provided
	if *out == "" {
		levelLower := strings.ToLower(*level)
		*out = fmt.Sprintf("docs/diagrams/asstBackend-%s.drawio", levelLower)
	}

	// Build diagram
	builder := &DiagramBuilder{
		Doc:    doc,
		Theme:  *theme,
		Layout: *layout,
	}

	nodes, edges := builder.Build()
	xml := GenerateMXGraph(nodes, edges)

	if *dry {
		fmt.Println("DRY RUN MODE")
		fmt.Println("============")
		fmt.Printf("Output: %s\n", *out)
		fmt.Printf("Level: %s\n", *level)
		fmt.Printf("Theme: %s\n", *theme)
		fmt.Printf("Layout: %s\n", *layout)
		fmt.Printf("Nodes: %d\n", len(nodes))
		fmt.Printf("Edges: %d\n", len(edges))
		fmt.Println("\nGenerated XML preview (first 500 chars):")
		if len(xml) > 500 {
			fmt.Println(xml[:500] + "...")
		} else {
			fmt.Println(xml)
		}
		return
	}

	// Ensure output directory exists
	outDir := filepath.Dir(*out)
	if err := os.MkdirAll(outDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory: %v\n", err)
		os.Exit(1)
	}

	// Write diagram file
	if err := os.WriteFile(*out, []byte(xml), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing diagram: %v\n", err)
		os.Exit(1)
	}

	// Write metadata
	assumptions := []string{
		"Diagram generated from C4 documentation",
		"Node positions may need manual adjustment for optimal layout",
		"Colors follow standard C4 conventions",
	}
	
	prompt := fmt.Sprintf("Generate %s diagram from %s", *level, *c4doc)
	if *c4doc == "" {
		prompt = fmt.Sprintf("Generate %s diagram", *level)
	}
	
	meta := GenerateMetadata(prompt, *level, *out, sources, assumptions)
	metaPath := *out + ".meta.md"
	if err := os.WriteFile(metaPath, []byte(meta), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to write metadata: %v\n", err)
	}

	fmt.Printf("✓ Generated %s diagram: %s\n", *level, *out)
	fmt.Printf("✓ Metadata: %s\n", metaPath)
	fmt.Printf("  Nodes: %d, Edges: %d\n", len(nodes), len(edges))
}
