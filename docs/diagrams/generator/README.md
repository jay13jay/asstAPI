# Diagram generator helper

This package provides a minimal generator and CLI used by the diagram agent. It includes:

- deterministic name normalization
- generation of a minimal mxGraph (draw.io) XML
- a tiny CLI for local dry-runs and writing `.drawio` files

Build:

```powershell
go build -o ./bin/diag-generator ./docs/diagrams/generator
```

Dry run example:

```powershell
./bin/diag-generator --prompt "L2 container: web, api, db" --dry-run=true
```
