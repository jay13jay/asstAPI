package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	out := flag.String("out", "docs/diagrams/asstBackend-l2-containers.drawio", "output .drawio file path")
	prompt := flag.String("prompt", "", "diagram prompt or path to C4 doc")
	dry := flag.Bool("dry-run", true, "dry run (print summary instead of writing)")
	flag.Parse()

	if *prompt == "" {
		fmt.Fprintln(os.Stderr, "missing --prompt")
		os.Exit(2)
	}

	// Minimal example: create two nodes and one edge
	apiID := NodeID("container", "API Server")
	dbID := NodeID("database", "Postgres")

	nodes := []Node{
		{ID: apiID, Label: "API Server", X: 120, Y: 60, W: 160, H: 60},
		{ID: dbID, Label: "Postgres", Style: "shape=cylinder;fillColor=#2b6cb0;fontColor=#ffffff;", X: 340, Y: 60, W: 120, H: 80},
	}
	edges := []Edge{{ID: EdgeID(apiID, dbID), Source: apiID, Target: dbID, Label: "SQL"}}

	xml := GenerateMXGraph(nodes, edges)

	if *dry {
		fmt.Println("DRY RUN: would write:", *out)
		fmt.Println("Generated preview (head):")
		if len(xml) > 400 {
			fmt.Println(xml[:400])
		} else {
			fmt.Println(xml)
		}
		return
	}

	if err := os.WriteFile(*out, []byte(xml), 0644); err != nil {
		fmt.Fprintln(os.Stderr, "write failed:", err)
		os.Exit(1)
	}

	// write metadata
	meta := fmt.Sprintf("prompt: %s\ngenerated: %s\nagent_version: v0.1.0\nmode: local\nsources: []\noutputs: [%s]\nassumptions: []\n", *prompt, time.Now().UTC().Format(time.RFC3339), *out)
	_ = os.WriteFile(*out+".meta.md", []byte(meta), 0644)

	fmt.Println("wrote", *out)
}
