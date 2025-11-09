# Diagram Helper Spec

Purpose

This file captures small, deterministic helper rules and examples the `diagram-creator` agent should use when generating draw.io (mxGraph) files. Keep this file under `.github/agents/` so the agent and humans can easily reference the conventions.

1) Deterministic ID format

- ID pattern: `node://<type>/<normalized-name>`
  - `<type>`: one of `system`, `container`, `component`, `actor`, `database`, `queue`, `external`
  - `<normalized-name>`: lowercased, whitespace and punctuation replaced with `-`, non-ascii removed where possible.
  - Examples:
    - `node://container/api-server`
    - `node://database/postgres`

Normalization pseudo-code (reference):

- trim input
- convert to NFC unicode
- lowercase
- remove characters not in [a-z0-9 \-_.]
- replace spaces and runs of punctuation with `-`
- collapse multiple `-` to single `-`
- strip leading/trailing `-`

2) Node attributes (canonical)

- id: deterministic ID (string)
- label: human friendly label (string)
- type: node type (one of the types above)
- link: optional URL string (must be sanitized)
- notes: optional longer text
- themeStyle: optional style token (filled from palettes)

3) Edge attributes

- id: `edge://<from-id>-to-<to-id>` (ensure length limits)
- label: optional (e.g., "HTTP", "gRPC", "SQL")
- style: `solid` or `dashed`
- dir: `->`, `<-`, `-` (directionality)

4) C4 → draw.io mapping (summary)

- System/Actor -> mxEllipse or person icon
- Container -> rounded rectangle
- Component -> rectangle with smaller subtitle text containing package path
- Database -> cylinder shape (draw.io has a "cylinder" style)
- Queue/Topic -> stacked rectangles or cloud-like icon
- External service -> rectangle with dashed border

5) Palettes (example tokens)

- light
  - background: #ffffff
  - primary: #0b5fff
  - accent: #00a676
  - db: #2b6cb0
  - text: #111827

- dark
  - background: #0b1220
  - primary: #7aa2ff
  - accent: #6ee7b7
  - db: #4c6ef5
  - text: #e6eef8

6) Minimal mxGraph XML example (2 nodes + edge)

This is a minimal, well-formed mxGraph fragment an agent can produce as the core of a `.drawio` file. A real file requires additional container attributes; this fragment shows node/edge structure the tests should assert for.

```xml
<?xml version="1.0" encoding="UTF-8"?>
<mxfile host="app.diagrams.net">
  <diagram id="diagram-1" name="Page-1">
    <mxGraphModel dx="1000" dy="800" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1">
      <root>
        <mxCell id="0"/>
        <mxCell id="1" parent="0"/>
        <!-- Node: API server -->
        <mxCell id="node://container/api-server" value="API Server" style="rounded=1;fillColor=#0b5fff;fontColor=#ffffff;" vertex="1" parent="1">
          <mxGeometry x="120" y="60" width="160" height="60" as="geometry"/>
        </mxCell>
        <!-- Node: Postgres -->
        <mxCell id="node://database/postgres" value="Postgres" style="shape=cylinder;fillColor=#2b6cb0;fontColor=#ffffff;" vertex="1" parent="1">
          <mxGeometry x="340" y="60" width="120" height="80" as="geometry"/>
        </mxCell>
        <!-- Edge -->
        <mxCell id="edge://node://container/api-server-to-node://database/postgres" style="edgeStyle=elbowEdgeStyle;rounded=0;" edge="1" parent="1" source="node://container/api-server" target="node://database/postgres">
          <mxGeometry relative="1" as="geometry"/>
        </mxCell>
      </root>
    </mxGraphModel>
  </diagram>
</mxfile>
```

7) `.meta.md` minimal template (created per diagram)

- prompt: "<original user prompt or issue body>"
- generated: 2025-11-09T12:34:56Z
- agent_version: v0.1.0
- mode: `local` | `github`
- sources: list of parsed files (paths)
- outputs: list of generated files
- assumptions: any mapping assumptions the agent made
- validation: `ok` | `xml-error: <message>`

8) Small validation guidance

- Verify XML is well-formed and contains the declared node IDs
- Ensure IDs match deterministic pattern
- If `auto-layout` is used and results in overlapping nodes, fallback to a simple horizontal layout and note in metadata

9) Example prompts (short)

- "L2 container: show web, api, db, worker. horizontal, light, name=asstBackend-l2-containers"
- "L3 chat: parse internal/handlers/chat.go and internal/service/chat.go; include component 'MessageRouter' and label edges with 'HTTP' and 'SQL'"

---

End of helper spec.
